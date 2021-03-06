package config

import (
	"errors"
	"fmt"
)

var (
	ErrMissingParam = errors.New("default values missing. Pass a map[string]interface{} to this function")
	ErrInvalidParam = errors.New("invalid param. map[string]interface{} expected")
)

type DBType uint32
const(
	DBPostgres = iota
	DBMySql
)

// Config holds application configuration
type Config struct {
	// GrpcAddress is the address that server listens to
	// Format: 127.0.0.1:9090
	GrpcAddress string

	// DBType represents type of database
	DBType	DBType

	// DBUsername database username
	DBUsername 	string

	// DBPassword database password
	DBPassword 	string

	// DBHostname database hostname.
	// Format: localhost | 127.0.0.1
	DBHostname	string

	// DBName database name
	DBName		string

	// DBPort database port
	DBPort		int
}

// Loader is an interface which is responsible for loading
// configs. We may have different implementations based on
// different environments and needs.
type Loader interface {
	// Load loads config from a source
	// params can be used for default values
	Load(params ...interface{}) (*Config, error)
}

// DefaultLoader loads default values
type DefaultLoader struct {}


func (dl *DefaultLoader) Load(params ...interface{}) (Config, error) {
	if len(params) == 0 {
		return Config{}, ErrMissingParam
	}

	p := params[0]
	m, ok := p.(map[string]interface{})
	if !ok {
		return Config{}, ErrInvalidParam
	}
	return Config{
		GrpcAddress: lookupStr(m, "GrpcAddress", "localhost:9090"),
		DBUsername:  lookupStr(m, "DBUsername", ""),
		DBPassword:  lookupStr(m, "DBPassword", ""),
		DBHostname:  lookupStr(m, "DBHostname", "localhost"),
		DBName:      lookupStr(m, "DBName", ""),
		DBPort:      lookupInt(m, "DBPort", 3306),
		DBType: 	 DBType(lookupInt(m, "DBType", DBPostgres)),
	}, nil
}

// lookup tries to find a val based on the given key. It returns default value if cannot
// find the key. Caller is responsible for type assertion
func lookup(m map[string]interface{}, key string, defaultVal interface{}) interface{} {
	if v, present := m[key]; present {
		return v
	}

	return defaultVal
}

// lookupInt calls lookup and converts its result to string.
func lookupStr(m map[string]interface{}, key string, defaultVal string) string {
	return fmt.Sprintf("%v", lookup(m, key, defaultVal))
}

// lookupInt calls lookup and converts its result to int.
// It returns defaultVal if type assertion is unsuccessful.
func lookupInt(m map[string]interface{}, key string, defaultVal int) int {
	v := lookup(m, key, defaultVal)
	i, ok := v.(int)
	if !ok {
		return defaultVal
	}

	return i
}