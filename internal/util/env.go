package util

import (
	"os"
)

func GetEnvVariables() map[string]string {
	return map[string]string{
		"PORT":     os.Getenv("PORT"),
		"HOST":     os.Getenv("HOST"),
		"APP_MODE": os.Getenv("APP_MODE"),
	}
}
