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

var LOCAL_PORT = GetEnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080")
var QUEUE_ACCOUNT_NAME = GetEnvOrDefault("QUEUE_ACCOUNT_NAME", "devstoreaccount1")
var QUEUE_ACCOUNT_KEY = GetEnvOrDefault("QUEUE_ACCOUNT_KEY", "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==")
var STAGE = GetEnvOrDefault("STAGE", "dev")
var MONGO_DB_CONNECTION_STRING = GetEnvOrDefault("MONGO_DB_CONNECTION_STRING", "mongodb://localhost:C2y6yDjf5%2FR%2Bob0N8A7Cgv30VRDJIWEHLM%2B4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw%2FJw%3D%3D@localhost:10255/admin?ssl=true")
