package deserialize

import (
	"gopkg.in/yaml.v2"
	gYaml "github.com/ghodss/yaml"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"bytes"
	"io"
)

//Deserializer deserializes.
type Deserializer struct {
	JSONUnmarshaler	*jsonpb.Unmarshaler
}

func New() *Deserializer {
	deserializer := &Deserializer{
		JSONUnmarshaler : &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		},
	}
	return deserializer
}

//YAMLToStruct populates your struct with specified byte array
func (d Deserializer) YAMLToStruct(data []byte, resp interface{}) error {
	err := yaml.Unmarshal(data, resp)
	return err
}

//YAMLToProto does a roundabout way of converting yaml to protobuf by using json as intermediary
func (d Deserializer) YAMLToProto(data []byte, msg proto.Message) error {
	jsonBytes, err := gYaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	d.JSONToProto(ioutil.NopCloser(bytes.NewReader(jsonBytes)), msg)
	return nil
}

//JSONToProto converts json stream to protobuf
func (d Deserializer) JSONToProto(requestBody io.ReadCloser, unmarshalObj proto.Message) (err error){
	defer requestBody.Close()
	if err := d.JSONUnmarshaler.Unmarshal(requestBody, unmarshalObj); err != nil {
		return err
	}
	return
}
