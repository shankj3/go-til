//HTTP related utility tools
package net

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/shankj3/go-til/log"
	"golang.org/x/oauth2/clientcredentials"
	"io/ioutil"
	"net/http"
	"net/url"
	"golang.org/x/oauth2"
)

//go:generate mockgen -source net.go -destination net.mock.go -package net

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

	//GetUrlResponse uses the OAuth Client to make an HTTP get call, and returns a normal response object. Caller is expected to close response body as per usual
	GetUrlResponse(url string) (*http.Response, error)

	//PostUrl will perform a post on the specified URL. It takes in a json formatted body
	//and returns an (optional) protobuf response
	PostUrl(url string, body string, unmarshalObj proto.Message) error

	// PostURLForm will post form data and return an http response
	PostUrlForm(url string, form url.Values) (*http.Response, error)

	// GetAuthClient will return the oauth authenticated client for more flexibility
	GetAuthClient() *http.Client
}

//OAuthClient is a client containing a pre-authenticated http client as returned by
//golang's oauth2 clientcredentials package as well as a protobuf json unmarshaler
type OAuthClient struct {
	AuthClient  *http.Client
	Unmarshaler jsonpb.Unmarshaler
}

type OAuthClientCreds interface {
	GetClientId() string
	GetClientSecret() string
	GetTokenURL() string
}


func (oc *OAuthClient) GetAuthClient() *http.Client {
	return oc.AuthClient
}

//Setup takes in OAuth2 credentials and returns a temporary token along with an error. uses two-legged OAuth model.
func (oc *OAuthClient) Setup(config OAuthClientCreds) (string, error) {
	var conf = clientcredentials.Config{
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
	oc.AuthClient = authClient
	return token.AccessToken, err
}

// SetupStaticToken will set up the OAuthClient to use a static token, will assume it will never expire.
//   Used for github api authentication, or any service that does not support two-legged OAuth
func (oc *OAuthClient) SetupStaticToken(config OAuthClientCreds) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetClientSecret()},
	)
	oc.AuthClient = oauth2.NewClient(ctx, ts)
	oc.Unmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	return config.GetClientSecret(), nil
}

// SetupNoAuthentication will set up OAuthClient to make calls with a regular *http.Client, all calls will be unauthenticated
func (oc *OAuthClient) SetupNoAuthentication() {
	oc.AuthClient = &http.Client{}
	oc.Unmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
}

//GetUrlResponse just uses the OAuth client to get the url.
func (oc *OAuthClient) GetUrlResponse(url string) (*http.Response, error) {
	return oc.AuthClient.Get(url)
}


func (oc *OAuthClient) GetUrl(url string, unmarshalObj proto.Message) error {
	// todo: this doesn't handle http response codes or anything... idk how much we need it in this case but seems weird
	resp, err := oc.AuthClient.Get(url)
	if err != nil {
		log.IncludeErrField(err).Error("can't get url ", url)
		return err
	}
	defer resp.Body.Close()
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

func (oc *OAuthClient) PostUrlForm(url string, form url.Values) (*http.Response, error) {
	return oc.AuthClient.PostForm(url, form)
}

func (oc *OAuthClient) PostUrl(url string, body string, unmarshalObj proto.Message) error {
	bodyBytes := []byte(body)
	resp, err := oc.AuthClient.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if resp.Body != nil {
		defer resp.Body.Close()
	}

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
