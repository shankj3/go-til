package nsqpb

import (
    "github.com/nsqio/go-nsq"

    "bitbucket.org/level11consulting/go-til/log"
)

// ProtoConsume wraps nsq.Message so that code outside the package can just add a UnmarshalProtoFunc
// that doesn't require messing with nsq fields. just write a function that unmarshals to your proto object
// and does work
// ...put in WORK.
type ProtoConsume struct {
	Handler      HandleMessage
    DecodeConfig *nsq.Config
    Config 		 *NsqConfig
    StopChan 	 chan int
}

// HandleMessage is an interface for unmarshalling your messages to a struct or protobuf message,
// then processing the object. Fulfilling this interface is how you would interact w/ the nsq channels
type HandleMessage interface {
	UnmarshalAndProcess([]byte) error
}

// NewProtoConsume returns a new ProtoConsume object with nsq configuration and
// nsqpb configuration
func NewProtoConsume() *ProtoConsume {
    config := nsq.NewConfig()
    return &ProtoConsume{
        DecodeConfig: config,
        Config:       DefaultNsqConf(),
    }
}

// NSQProtoConsume is a wrapper for `p.Handler.UnmarshalAndProcess` --> `nsq.HandlerFunc`
func (p *ProtoConsume) NSQProtoConsume(msg *nsq.Message) error {
	log.Log().Debug("Inside wrapper for UnmarshalAndProcess")
    if err := p.Handler.UnmarshalAndProcess(msg.Body); err != nil {
        log.IncludeErrField(err).Warn("nsq proto consume error")
        return err
    }
    return nil
}

// Consume messages on a given topic / channel in NSQ protoconsume's UnmarshalProtoFunc will be added with
// a wrapper as a handler for the consumer. The ip address of the NSQLookupd instance
// can be set by the environment variable NSQLOOKUPD_IP, but will default to 127.0.0.1
func (p *ProtoConsume) ConsumeMessages(topicName string, channelName string) error {
    log.Log().Debug("Inside Consume Messages")
    c, err := nsq.NewConsumer(topicName, channelName, p.DecodeConfig)
    if err != nil {
        log.IncludeErrField(err).Warn("cannot create nsq consumer")
        return err
    }
	p.StopChan = c.StopChan
    c.SetLogger(NSQLogger{}, nsq.LogLevelError)
    c.AddHandler(nsq.HandlerFunc(p.NSQProtoConsume))

    if err = c.ConnectToNSQLookupd(p.Config.LookupDAddress()); err != nil {
        log.IncludeErrField(err).Warn("cannot connect to nsq")
        return err
    }
    return nil
}