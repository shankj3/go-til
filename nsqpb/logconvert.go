package nsqpb

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"strings"
	"bitbucket.org/level11consulting/go-til/log"
)

func ConvertLogLevel(level logrus.Level) nsq.LogLevel {
	switch level {
	case logrus.DebugLevel:
		return nsq.LogLevelDebug
	case logrus.InfoLevel:
		return nsq.LogLevelInfo
	case logrus.WarnLevel:
		return nsq.LogLevelWarning
	case logrus.ErrorLevel:
		return nsq.LogLevelError
	}
	return nsq.LogLevelWarning
}

var (
	nsqDebugLevel = nsq.LogLevelDebug.String()
	nsqInfoLevel  = nsq.LogLevelInfo.String()
	nsqWarnLevel  = nsq.LogLevelWarning.String()
	nsqErrLevel   = nsq.LogLevelError.String()
)

type NSQLogger struct{}

func NewNSQLogger() (logger NSQLogger, level nsq.LogLevel) {
	return NewNSQLoggerAtLevel(log.GetLogLevel())
}

func NewNSQLoggerAtLevel(lvl logrus.Level) (logger NSQLogger, level nsq.LogLevel){
	logger = NSQLogger{}
	// hard coding this becuase nsq debug is so annoying
	level = ConvertLogLevel(logrus.InfoLevel)
	return
}

func (n NSQLogger) Output(_ int, s string) error {
	if len(s) > 3 {
		msg := strings.TrimSpace(s[3:])
		switch s[:3] {
		case nsqDebugLevel:
			log.Log().Debugln(msg)
		case nsqInfoLevel:
			log.Log().Infoln(msg)
		case nsqWarnLevel:
			log.Log().Warnln(msg)
		case nsqErrLevel:
			log.Log().Errorln(msg)
		default:
			log.Log().Infoln(msg)
		}
	}
	return nil
}
