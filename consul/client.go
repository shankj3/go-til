package consul

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/hashicorp/consul/api"
)

var (
	consulCalls = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "consul_requests_total",
		Help: "all consul requests made",
		// status: fail | success
	}, []string{"status"})
)

func record(err error) {
	var status string
	if err == nil {
		status = "success"
	} else {
		status = "fail"
	}
	consulCalls.WithLabelValues(status).Inc()
}

func init() {
	prometheus.MustRegister(consulCalls)
}

type Consuletty interface {
	AddKeyValue(key string, value []byte) error
	RemoveValue(key string) error
	RemoveValues(prefix string) error
	GetKeys(prefix string) (keys []string, err error)
	GetKeyValue(key string) (*api.KVPair, error)
	GetKeyValues(prefix string) (api.KVPairs, error)
	RegisterService(addr string, port int, name string) error
	RemoveService(name string) error
	IsConnected() bool
}

//Consulet is a wrapper for interfacing with consul
type Consulet struct {
	Client *api.Client
	Config *api.Config
	Connected bool
}

//Default will assume consul is running at localhost:8500
func Default() (*Consulet, error) {
	consulet := &Consulet{}
	consulet.Config = api.DefaultConfig()
	c, err := api.NewClient(consulet.Config)

	if err != nil {
		return nil, err
	}
	consulet.Client = c
	consulet.checkIfConnected()
	return consulet, nil
}

//New allows for configuration of consul host + port
func New(consulHost string, consulPort int) (*Consulet, error) {
	consulet := &Consulet{}

	consulet.Config = &api.Config{
		Address: consulHost + ":" + strconv.Itoa(consulPort),
	}
	c, err := api.NewClient(consulet.Config)

	if err != nil {
		return nil, err
	}

	consulet.Client = c
	consulet.checkIfConnected()
	return consulet, nil
}

func (consul *Consulet) IsConnected() bool {
	return consul.Connected
}

//RegisterService registers a service at specified host, port, with name
func (consul *Consulet) RegisterService(addr string, port int, name string) (err error) {
	reg := &api.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	err = consul.Client.Agent().ServiceRegister(reg)
	consul.updateConnection(err)
	return
}

//RemoveService removes a service by name
func (consul *Consulet) RemoveService(name string) error {
	err := consul.Client.Agent().ServiceDeregister(name)
	consul.updateConnection(err)
	return err
}

//TODO: should key value operations be atomic??? Can switch to use CAS
func (consul *Consulet) AddKeyValue(key string, value []byte) (err error) {
	defer record(err)
	kv := consul.Client.KV()
	kvPair := &api.KVPair{
		Key:   key,
		Value: value,
	}
	_, err = kv.Put(kvPair, nil)
	consul.updateConnection(err)
	return err
}

//RemoveValue removes value at specified key
func (consul *Consulet) RemoveValue(key string) (err error) {
	defer record(err)
	kv := consul.Client.KV()
	_, err = kv.Delete(key, nil)
	consul.updateConnection(err)
	return err
}

// Remove values at specified prefix (like `consul kv delete -recurse /prefix`)
func (consul *Consulet) RemoveValues(prefix string) (err error) {
	defer record(err)
	kv := consul.Client.KV()
	_, err = kv.DeleteTree(prefix, nil)
	consul.updateConnection(err)
	return err
}

// GetKeys uses consul default of separator ("/")
func (consul *Consulet) GetKeys(prefix string) (keys []string, err error) {
	defer record(err)
	kv := consul.Client.KV()
	keys, _, err = kv.Keys(prefix, "", nil)
	if err != nil {
		return
	}
	return
}

//GetKeyValue gets key/value at specified key
func (consul *Consulet) GetKeyValue(key string) (pair *api.KVPair, err error) {
	defer record(err)
	kv := consul.Client.KV()
	pair, _, err = kv.Get(key, nil)
	consul.updateConnection(err)
	return pair, err
}

//GetKeyValue gets key/value list at specified prefix
func (consul *Consulet) GetKeyValues(prefix string) (pairs api.KVPairs, err error) {
	kv := consul.Client.KV()
	pairs, _, err = kv.List(prefix, nil)
	consul.updateConnection(err)
	return pairs, err
}

func (consul *Consulet) CreateNewSemaphore(path string, limit int) (sema *api.Semaphore, err error) {
	defer record(err)
	sessionName := fmt.Sprintf("semaphore_%s", path)
	// create new session. the health check is just gossip failure detector, session will
	// be held as long as the default serf health check hasn't declared node unhealthy.
	// if that node is unhealthy, it probably won't be able to finish running the build so someone
	// else can pick it up... sidenote.. we need to handle if a worker goes down.
	sessionId, _, err := consul.Client.Session().Create(&api.SessionEntry{
		Name: sessionName,
	}, nil)
	if err != nil {
		consul.updateConnection(err)
		return nil, err
	}
	semaphoreOpts := &api.SemaphoreOptions{
		Prefix:      path,
		Limit:       limit,
		Session:     sessionId,
		SessionName: sessionName,
	}
	sema, err = consul.Client.SemaphoreOpts(semaphoreOpts)
	if err != nil {
		consul.updateConnection(err)
		return nil, err
	}
	return sema, nil
}

//checkIfConnected is called inside of initilization functions to properly update
//the Connected bool flag. It just tries to read at a key that shouldn't exist
func (consul *Consulet) checkIfConnected() {
	consul.GetKeyValue("connection_test")
}

//updateConnection takes in an error message and will update the Connection bool to be false if we get connection refused error.
//Also logs error if exists. TODO? Can add a new field to Consulet struct showing most recent failure's error message
func (consul *Consulet) updateConnection(err error) {
	//assumes that
	if err == nil || (err != nil && !strings.Contains(err.Error(), ": connection refused")) {
		consul.Connected = true
	} else {
		consul.Connected = false
		// connection errors are happening
		fmt.Println("changing connected to false, err: ", err.Error())
	}
}