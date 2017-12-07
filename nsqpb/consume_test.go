package nsqpb

import (
	"testing"
	"bitbucket.org/level11consulting/go-til/test"
)

func TestProtoConsume_Topics(t *testing.T) {
	protoConsumer := NewProtoConsume()
	protoConsumer.AddTopic("marianne1")
	protoConsumer.AddTopic("marianne2")
	if len(protoConsumer.GetTopics()) != 2 {
		t.Error(test.GenericStrFormatErrors("number of supported topics", 2, len(protoConsumer.GetTopics())))
	}
	if protoConsumer.GetTopics()[0] != "marianne1" {
		t.Error(test.GenericStrFormatErrors("first topic", "marianne1", protoConsumer.GetTopics()[0]))
	}
	if protoConsumer.GetTopics()[1] != "marianne2" {
		t.Error(test.GenericStrFormatErrors("second topic", "marianne2", protoConsumer.GetTopics()[1]))
	}

	protoConsumer.DeleteTopic("marianne2")
	if len(protoConsumer.GetTopics()) != 1 {
		t.Error(test.GenericStrFormatErrors("number of supported topics", 1, len(protoConsumer.GetTopics())))
	}
	if protoConsumer.GetTopics()[0] != "marianne1" {
		t.Error(test.GenericStrFormatErrors("first topic", "marianne1", protoConsumer.GetTopics()[0]))
	}
}