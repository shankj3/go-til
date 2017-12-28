package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
	"strings"
)

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
func (consul *Consulet) AddKeyValue(key string, value []byte) error {
	kv := consul.Client.KV()
	kvPair := &api.KVPair{
		Key:   key,
		Value: value,
	}
	_, err := kv.Put(kvPair, nil)
	consul.updateConnection(err)
	return err
}

//RemoveValue removes value at specified key
func (consul *Consulet) RemoveValue(key string) error {
	kv := consul.Client.KV()
	_, err := kv.Delete(key, nil)
	consul.updateConnection(err)
	return err
}

// Remove values at specified prefix (like `consul kv delete -recurse /prefix`)
func (consul *Consulet) RemoveValues(prefix string) error {
	kv := consul.Client.KV()
	_, err := kv.DeleteTree(prefix, nil)
	consul.updateConnection(err)
	return err
}

//GetKeyValue gets key/value at specified key
func (consul *Consulet) GetKeyValue(key string) (*api.KVPair, error) {
	kv := consul.Client.KV()
	val, _, err := kv.Get(key, nil)
	consul.updateConnection(err)
	return val, err
}

//GetKeyValue gets key/value list at specified prefix
func (consul *Consulet) GetKeyValues(prefix string) (api.KVPairs, error) {
	kv := consul.Client.KV()
	val, _, err := kv.List(prefix, nil)
	consul.updateConnection(err)
	return val, err
}

func (consul *Consulet) CreateNewSemaphore(path string, limit int) (*api.Semaphore, error) {
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
	sema, err := consul.Client.SemaphoreOpts(semaphoreOpts)
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
	if err == nil || (err != nil && !strings.Contains(err.Error(), "getsockopt: connection refused")) {
		consul.Connected = true
	} else {
		consul.Connected = false
		// connection errors are happening
		fmt.Println("changing connected to false, err: ", err.Error())
	}
}