package config

import "os"

type Stage string

const (
	DEV     = "dev"
	STAGING = "stg"
	PROD    = "prod"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var LOCAL_PORT = GetEnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "7071")
var QUEUE_ACCOUNT_NAME = GetEnvOrDefault("QUEUE_ACCOUNT_NAME", "devstoreaccount1")
var QUEUE_ACCOUNT_KEY = GetEnvOrDefault("QUEUE_ACCOUNT_KEY", "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==")
var STAGE = GetEnvOrDefault("STAGE", "dev")
