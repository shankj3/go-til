package nsqpb

import (
    "fmt"
    "github.com/nsqio/go-nsq"
    "github.com/shankj3/go-til/log"
    "os"
    "runtime/debug"
    "time"
)

// ProtoConsume wraps nsq.Message so that code outside the package can just add a UnmarshalProtoFunc
// that doesn't require messing with nsq fields. just write a function that unmarshals to your proto object
// and does work
// ...put in WORK.
type ProtoConsume struct {
    Handler          HandleMessage
    DecodeConfig     *nsq.Config
    Config           *NsqConfig
    StopChan         chan int
    ConsumerRecovery func()
    MessageRecovery  func(message *nsq.Message)
    consumers        []*nsq.Consumer
    topics           []string
}

// HandleMessage is an interface for unmarshalling your messages to a struct or protobuf message,
// then processing the object. Fulfilling this interface is how you would interact w/ the nsq channels
type HandleMessage interface {
    UnmarshalAndProcess(msg []byte, done chan int, finish chan int) error
}

func defaultMsgRecovery(message *nsq.Message) {
    if r := recover(); r != nil {
        fmt.Println("OOOHHH MAN, A PANIC HAPPENED!!")
        debug.PrintStack()
        fmt.Println("going to try to requeue")
        message.Requeue(0)
        os.Exit(1)
    }

}

func defaultConsumerRecovery() {
    if r := recover(); r != nil {
        fmt.Println("OOOHHH MAN, A PANIC HAPPENED!!")
        debug.PrintStack()
        os.Exit(1)
    }
}

// NewDefaultProtoConsume returns a new ProtoConsume object with nsq configuration and
// nsqpb configuration. also sets default message recovery and consumer recovery functions
func NewDefaultProtoConsume() *ProtoConsume {
    config := nsq.NewConfig()
    nsqpbConf := DefaultNsqConf()
    config.MsgTimeout = time.Second * time.Duration(nsqpbConf.Timeout)
    return &ProtoConsume{
        DecodeConfig:     config,
        Config:           nsqpbConf,
        MessageRecovery:  defaultMsgRecovery,
        ConsumerRecovery: defaultConsumerRecovery,
    }
}

// NSQProtoConsume is a wrapper for `p.Handler.UnmarshalAndProcess` --> `nsq.HandlerFunc`
func (p *ProtoConsume) NSQProtoConsume(msg *nsq.Message) error {
    log.Log().WithField("nsqMsgId", string(msg.ID[:])).Info("receiving nsq proto msg")
    defer p.MessageRecovery(msg)
    done := make(chan int)
    finish := make(chan int)
    log.Log().Debug("Inside wrapper for UnmarshalAndProcess")
    go p.Handler.UnmarshalAndProcess(msg.Body, done, finish)
    // TODO: error for requeing? quizas?
    for {
        select {
        case <-done:
            log.Log().WithField("nsqMsgId", string(msg.ID[:])).Info("received on done channel, will stop sending TOUCH commands to nsq")
            msg.Finish()
            return nil
        case <-finish:
            log.Log().WithField("nsqMsgId", string(msg.ID[:])).Info("recieved on finish channel, calling msg.Finish()")
            msg.Finish()
            return nil
        default:
            msg.Touch()
            time.Sleep(time.Second * time.Duration(p.Config.TouchInterval))
        }
        for _, consumer := range p.consumers {
            //if consumer.IsStarved() {
            //  log.Log().Error("the consumer is starved!!")
            //}
            stats := consumer.Stats()
            log.Log().WithField("connections", fmt.Sprintf("%d", stats.Connections)).
                WithField("messagesReceived", fmt.Sprintf("%d", stats.MessagesReceived)).
                WithField("messagesFinished", fmt.Sprintf("%d", stats.MessagesFinished)).
                WithField("messagesRequeued", fmt.Sprintf("%d", stats.MessagesRequeued)).Debug("consumer stats")
        }
    }
}

func (p *ProtoConsume) GetStats() []*nsq.ConsumerStats {
    var stats []*nsq.ConsumerStats
    for _, consumer := range p.consumers {
        stats = append(stats, consumer.Stats())
    }
    return stats
}

// Consume messages on a given topic / channel in NSQ protoconsume's UnmarshalProtoFunc will be added with
// a wrapper as a handler for the consumer. The ip address of the NSQLookupd instance
// can be set by the environment variable NSQLOOKUPD_IP, but will default to 127.0.0.1
func (p *ProtoConsume) ConsumeMessages(topicName string, channelName string) error {
    defer p.ConsumerRecovery()
    log.Log().Debug("Inside Consume Messages")
    c, err := nsq.NewConsumer(topicName, channelName, p.DecodeConfig)
    if err != nil {
        log.IncludeErrField(err).Warn("cannot create nsq consumer")
        return err
    }
    log.Log().Debugf("Changing max in flight to %d", 3)
    c.ChangeMaxInFlight(3)
    p.StopChan = c.StopChan
    c.SetLogger(NSQLogger{}, nsq.LogLevelError)
    //c.AddHandler(nsq.HandlerFunc(p.NSQProtoConsume))
    c.AddConcurrentHandlers(nsq.HandlerFunc(p.NSQProtoConsume), 2)
    p.consumers = append(p.consumers, c)
    if err = c.ConnectToNSQLookupd(p.Config.LookupDAddress()); err != nil {
        log.IncludeErrField(err).Warn("cannot connect to nsq")
        return err
    }
    return nil
}

func (p *ProtoConsume) Pause() {
    for _, consumer := range p.consumers {
        consumer.ChangeMaxInFlight(0)
    }
}

func (p *ProtoConsume) UnPause() {
    for _, consumer := range p.consumers {
        consumer.ChangeMaxInFlight(p.Config.MaxInFlight + 3)
    }
}

// Adds a supported topic to store on consumer
func (p *ProtoConsume) AddTopic(supportedTopic string) {
    p.topics = append(p.topics, supportedTopic)
}

// Retrieves all consumer supported topics
func (p *ProtoConsume) GetTopics() []string {
    return p.topics
}

//TODO: does it matter to add a bool for pass/fail
func (p *ProtoConsume) DeleteTopic(toRemove string) {
    for i, t := range p.topics {
        if t == toRemove {
            p.topics = append(p.topics[:i], p.topics[i+1:]...)
            break
        }
    }
}
