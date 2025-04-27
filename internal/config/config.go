package config

import (
	"github.com/lubie-koty/rpc-compute-service-simple/internal/validate"
)

type Config struct {
	Host string
	Port int
}

var AppConfig Config

func InitConfig(envVars map[string]string) error {
	portValue, err := validate.ValidatePort(envVars["PORT"])
	if err != nil {
		return err
	}
	hostValue, err := validate.ValidateHost(envVars["HOST"])
	if err != nil {
		return err
	}
	AppConfig = Config{
		Host: hostValue,
		Port: portValue,
	}
	return nil
}
