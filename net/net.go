//HTTP related utility tools
package net

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"golang.org/x/oauth2/clientcredentials"
	"bitbucket.org/level11consulting/go-til/log"
)

//TODO: what happens if I never create a instance of logger to hold on to?

var (
	FileNotFound = errors.New("could not find raw data at url")
)

//HttpClient is an http client interface that you can implement.
type HttpClient interface {
	//GetUrl will perform a GET on the specified URL and return the appropriate protobuf response
	GetUrl(url string, unmarshalObj proto.Message) error

	//GetUrlRawData will return raw data at specified URL in a byte array
	GetUrlRawData(url string) ([]byte, error)

	//PostUrl will perform a post on the specified URL. It takes in a json formatted body
	//and returns an (optional) protobuf response
	PostUrl(url string, body string, unmarshalObj proto.Message) error
}


//OAuthClient is a client containing a pre-authenticated http client as returned by
//golang's oauth2 clientcredentials package as well as a protobuf json unmarshaler
type OAuthClient struct {
	AuthClient  http.Client
	Unmarshaler jsonpb.Unmarshaler
}

type OAuthClientCreds interface {
	GetClientId() string
	GetClientSecret() string
	GetTokenURL() string
}

//Setup takes in OAuth2 credentials and returns a temporary token along with an error
func (oc *OAuthClient) Setup(config OAuthClientCreds) (string, error){
	var conf = clientcredentials.Config {
		ClientID:     config.GetClientId(),
		ClientSecret: config.GetClientSecret(),
		TokenURL:     config.GetTokenURL(),
	}
	var ctx = context.Background()
	token, err := conf.Token(ctx)
	if err != nil {
		log.IncludeErrField(err).Error("Unable to retrieve token for " + config.GetClientId() + " at " + config.GetTokenURL())
		return "", err
	}

	authClient := conf.Client(ctx)

	oc.Unmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	oc.AuthClient = *authClient
	return token.AccessToken, err
}


func (oc *OAuthClient) GetUrl(url string, unmarshalObj proto.Message) error {
	resp, err := oc.AuthClient.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.IncludeErrField(err).Error("can't get url ", url)
		return err
	}
	reader := bufio.NewReader(resp.Body)

	if err := oc.Unmarshaler.Unmarshal(reader, unmarshalObj); err != nil {
		log.IncludeErrField(err).Error("failed to parse response from ", url)
		return err
	}

	return nil
}


func (oc *OAuthClient) GetUrlRawData(url string) ([]byte, error) {
	resp, err := oc.AuthClient.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.IncludeErrField(err).Error("can't get url ", url)
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, FileNotFound
	}
	bytez, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytez, nil
}


func (oc *OAuthClient) PostUrl(url string, body string, unmarshalObj proto.Message) error {
	bodyBytes := []byte(body)
	resp, err := oc.AuthClient.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	defer resp.Body.Close()

	if err != nil {
		log.IncludeErrField(err).Error("can't post to url ", url)
		return err
	}

	if unmarshalObj != nil {
		reader := bufio.NewReader(resp.Body)

		if err := oc.Unmarshaler.Unmarshal(reader, unmarshalObj); err != nil {
			log.IncludeErrField(err).Error("failed to parse response from ", url)
			return err
		}
	} else {
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Log().Debug(string(respBody))
	}

	return nil
}
