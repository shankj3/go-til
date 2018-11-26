package vault

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/vault/api"
	ocelog "github.com/shankj3/go-til/log"
)

//go:generate mockgen -source client.go -destination client.mock.go -package vault

// VaultCIPath is the base path for vault. Will be formatted to include the user or group when
// setting or retrieving credentials.
var VaultCIPath = "secret/data/%s"

//var Token = "e57369ad-9419-cc03-9354-fc227b06f795"

// Some blog said that changing any *api.Client functions to take in a n interface instead
// will make testing easier. I agree, just have to figure out how to do this properly without
// wasting memory

//type ApiClient interface {
//	Logical() *ApiLo
//
//}
//
//type ApiLogical interface {
//	Read(path string) (*api.Secret, error)
//	Write(path string, data map[string]interface{}) (*api.Secret, error)
//}

// Vaulty is the go-til wrapper interface to the Vault API
type Vaulty interface {
	AddUserAuthData(user string, data map[string]interface{}) (*api.Secret, error)
	GetUserAuthData(user string) (map[string]interface{}, error)
	AddVaultData(path string, data map[string]interface{}) (*api.Secret, error)
	GetVaultData(path string) (map[string]interface{}, error)
	GetVaultSecret(path string) (*api.Secret, error)
	CreateToken(request *api.TokenCreateRequest) (token string, err error)
	CreateThrowawayToken() (token string, err error)
	CreateVaultPolicy() error
	GetAddress() string
	Healthy() bool
	DeletePath(path string) error
	RenewLeaseForever(secret *api.Secret) error
	RenewLeaseOnce(leaseID string, increment int) (*api.Secret, error)
	//EnableDatabaseSecretEngine(config map[string]interface{}) error
	//DisableDatabaseSecretEngine() error
}

// VaultyImpl is the go-til wrapper to the Vault client
type VaultyImpl struct {
	Client *api.Client
	Config *api.Config
}

// GetInitVault will return an authenticated Vault client
// Use this function as a singleton essentially.
// todo,  flesh out docs, for now look at hookhandler for use.
func GetInitVault(once sync.Once, vaultCached Vaulty) (Vaulty, error) {
	once.Do(func() {
		ocev, err := NewEnvAuthClient()
		if err != nil {
			return
		}
		vaultCached = ocev
	})
	return vaultCached, nil
}

// NewEnvAuthClient will set the Client token based on the environment variable `$VAULT_TOKEN`.
// Will return error if it is not set. Returns configured ocevault struct
func NewEnvAuthClient() (Vaulty, error) {
	var token string
	if token = os.Getenv("VAULT_TOKEN"); token == "" {
		return &VaultyImpl{}, errors.New("$VAULT_TOKEN not set")
	}
	return NewAuthedClient(token)
}

// NewAuthedClient will return a client with default configurations and the Token attached to it.
// Vault URL configured through VAULT_ADDR environment variable.
func NewAuthedClient(token string) (val Vaulty, err error) {
	valImpl := &VaultyImpl{}
	valImpl.Config = api.DefaultConfig()
	valImpl.Client, err = api.NewClient(valImpl.Config)
	if err != nil {
		return valImpl, nil
	}
	valImpl.Client.SetToken(token)
	// this action is idempotent, and since we *need* this policy for generating tokens, might as well?
	// i guess?
	valImpl.CreateVaultPolicy()
	return valImpl, nil
}

// RenewToken is a wrapper to the Vault api. Renews the token for 24 hours.
func (val *VaultyImpl) RenewToken() error {
	_, err := val.Client.Auth().Token().RenewSelf(86400)
	return err
}

// RenewLeaseForever is intended to be run as a goroutine. Will wait for 75% of ttl (secret.LeaseDuration), then try to renew the secret with same ttl
func (val *VaultyImpl) RenewLeaseForever(secret *api.Secret) error {
	currentSecret := secret
	err := errors.New("The RenewLeaseForever loop failed before setting an error. This shouldn't ever happen")
	for {
		time.Sleep(time.Duration((currentSecret.LeaseDuration/4)*3) * time.Second)
		//time.Sleep(time.Duration(60) * time.Second)
		ocelog.Log().Debugf("About to renew the lease for %s", secret.LeaseID)
		currentSecret, err = val.RenewLeaseOnce(currentSecret.LeaseID, currentSecret.LeaseDuration)
		if err != nil {
			return err
		}
		//ocelog.Log().Debugf("Current secret: %v", currentSecret)
	}
	// This should be unreachable
	//return err
}

// RenewLeaseOnce is a wrapper to the Vault API secret renew
func (val *VaultyImpl) RenewLeaseOnce(leaseID string, increment int) (*api.Secret, error) {
	return val.Client.Sys().Renew(leaseID, increment)
}

// Healthy returns true if the Vault server returns a HealthResponse. Otherwise returns false.
func (val *VaultyImpl) Healthy() bool {
	_, err := val.Client.Sys().Health()
	if err == nil {
		return true
	}
	return false
}

// AddUserAuthData will add the values of the data map to the path of the CI user creds
// CI vault path set off of base path VaultCIPath
func (val *VaultyImpl) AddUserAuthData(user string, data map[string]interface{}) (*api.Secret, error) {
	return val.Client.Logical().Write(fmt.Sprintf(VaultCIPath, user), data)
}

// GetAddress returns the Vault client address
func (val *VaultyImpl) GetAddress() string {
	return val.Client.Address()
}

// AddVaultData will add the values of the data map to the path of the CI user creds
// CI vault path set off of base path VaultCIPath
func (val *VaultyImpl) AddVaultData(path string, data map[string]interface{}) (*api.Secret, error) {
	return val.Client.Logical().Write(path, data)
}

// GetVaultData Reads from a given Vault path, but only returns the Data element
func (val *VaultyImpl) GetVaultData(path string) (map[string]interface{}, error) {
	secret, err := val.Client.Logical().Read(path)
	ocelog.Log().Debugf("Getting secret at path %s", path)
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, NotFound(fmt.Sprintf("user data not found, path searched: %s", path))
	}
	return secret.Data, nil
}

// GetVaultSecret Reads from a given Vault path. It is a lazy copy/paste of GetVaultData, but instead returns the full secret
func (val *VaultyImpl) GetVaultSecret(path string) (*api.Secret, error) {
	secret, err := val.Client.Logical().Read(path)
	ocelog.Log().Debugf("Getting secret at path %s", path)
	if err != nil {
		return nil, err
	}
	if secret == nil {
		return nil, NotFound(fmt.Sprintf("user data not found, path searched: %s", path))
	}
	return secret, nil
}

// GetUserAuthData will return the Data attribute of the secret you get at the path of the CI user creds, ie all the
// key-value fields that were set on it
func (val *VaultyImpl) GetUserAuthData(user string) (map[string]interface{}, error) {
	path := fmt.Sprintf(VaultCIPath, user)
	return val.GetVaultData(path)
}

// DeletePath will format the path with prepending our mount path (secret/data) and then deleting at the fully qualified path
// will return any errors from the Vault API
func (val *VaultyImpl) DeletePath(path string) error {
	fullPath := fmt.Sprintf(VaultCIPath, path)
	_, err := val.Client.Logical().Delete(fullPath)
	return err
}

// CreateToken creates an Auth token using the val.Client's creds. Look at api.TokenCreateRequest docs
// for how to configure the token. Will return any errors from the create request.
func (val *VaultyImpl) CreateToken(request *api.TokenCreateRequest) (token string, err error) {
	secret, err := val.Client.Auth().Token().Create(request)
	if err != nil {
		return
	}
	token = secret.Auth.ClientToken
	return
}

// CreateThrowawayToken creates a single use token w/ same privileges as client.
// *single use* really means enough uses to initialize the client and make one call to actually
// get data
// todo: add ocevault policy for reading the secrets/ci/user path
func (val *VaultyImpl) CreateThrowawayToken() (token string, err error) {
	tokenReq := &api.TokenCreateRequest{
		//Policies: 		[]string{"ocevault"}, // todo: figure out why this doesn't work...
		TTL:     "30m",
		NumUses: 3,
	}
	//oce.Client.Auth().Token().Create(&api.})
	return val.CreateToken(tokenReq)
}

// CreateVaultPolicy creates a policy for r/w ops on only the path that credentials are on, which is `secret/ci/creds`.
// Tokens that are one-off and passed to the workers for building will get this access.
func (val *VaultyImpl) CreateVaultPolicy() error {
	err := val.Client.Sys().PutPolicy("ocevault", "path \"secret/ci/creds\" {\n capabilities = [\"read\", \"list\"]\n}")
	if err != nil {
		return err
	}
	return nil
}

// NotFound returns a ErrNotFound string wrapper
func NotFound(msg string) *ErrNotFound {
	return &ErrNotFound{msg: msg}
}

// ErrNotFound is a string wrapping error type
type ErrNotFound struct {
	msg string
}

// Error returns the error message from ErrNotFound struct
func (e *ErrNotFound) Error() string {
	return e.msg
}

//// FIXME: We are not able to write tests that use Database Secret Engine until we can enable the backend in a test context.
//// EnableDatabaseSecretEngine will enable the Database Secret Engine via the Vault api
//func (val *VaultyImpl) EnableDatabaseSecretEngine(config map[string]interface{}) error {
//
//	mountInput := api.MountInput{}
//
//	//return secret, nil
//	return val.Client.Sys().Mount("database/config/ocelot", &mountInput)
//}
//
//// FIXME
//// DisableDatabaseSecretEngine will disable the Database Secret Engine via the Vault api
//func (val *VaultyImpl) DisableDatabaseSecretEngine() error {
//	return nil
//}
