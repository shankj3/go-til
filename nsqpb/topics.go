package nsqpb

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/shankj3/ocelot/protos/out"
//	"net/http"
//)
//
//const (
//	PRTopic   = "pr"
//	PushTopic = "push"
//
//)
//
////TODO: put this somewhere else - this needs to live in ocelot
//
//var SupportedTopics = [2]string{PushTopic, PRTopic}
//
//// extends proto.Message interface
//type BundleProtoMessage interface {
//	GetCheckoutHash() string
//	Reset()
//	String() string
//	ProtoMessage()
//}
//
//
//
//func TopicsUnmarshalObj(topic string) BundleProtoMessage {
//	switch topic {
//	case PRTopic:   return &protos.PRBuildBundle{}
//	case PushTopic: return &protos.PushBuildBundle{}
//	default:        return nil
//	}
//}
//
//type Topics struct {
//	topics []string
//}
//
//func GetTopics(nsqdlookupHostPort string) (*Topics, error) {
//	nsqdlookupAddr := fmt.Sprintf("http://%s/topics", nsqdlookupHostPort)
//	resp, err := http.Get(nsqdlookupAddr)
//	if err != nil {
//		return nil, err
//	}
//	tops := &Topics{}
//	defer resp.Body.Close()
//	if err = json.NewDecoder(resp.Body).Decode(tops); err != nil {
//		return nil, err
//	}
//	return tops, nil
//}
//