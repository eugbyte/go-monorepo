package config

import (
	configlib "github.com/eugbyte/monorepo/libs/config"
)

type vars struct {
	LOCAL_PORT                 string
	QUEUE_ACCOUNT_NAME         string
	QUEUE_ACCOUNT_KEY          string
	MONGO_DB_CONNECTION_STRING string
	VAULT_URI                  string
	NOTIFY_BASE_URL            string
	VAPID_PRIVATE_KEY          string
	VAPID_PUBLIC_KEY           string
	VAPID_EMAIL                string
}

// redeclare variables to avoid confusion between current config pkg and configlib pkg
type STAGE = configlib.STAGE
type FetchVal = configlib.FetchVal

var DEV = configlib.DEV
var STAGING = configlib.STAGING
var PROD = configlib.PROD

var Stage func() configlib.STAGE = configlib.Stage
var EnvOrDefault func(key string, defaultValue string) string = configlib.EnvOrDefault
var QueueBaseURL func(stage STAGE, accountName string) string = configlib.QueueBaseURL
var FetchAll func(fetchVal FetchVal, secretNames ...string) ([]string, error) = configlib.FetchAll

func New() vars {
	var env_vars map[STAGE]vars = map[STAGE]vars{
		DEV: {
			LOCAL_PORT:                 EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
			QUEUE_ACCOUNT_NAME:         "devstoreaccount1",
			QUEUE_ACCOUNT_KEY:          "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==",
			MONGO_DB_CONNECTION_STRING: EnvOrDefault("MONGO_DB_CONNECTION_STRING", ""),
			VAULT_URI:                  "https://kv-notify-secrets-stg-ea.vault.azure.net",
			NOTIFY_BASE_URL:            "http://localhost:7071/api",
			VAPID_PRIVATE_KEY:          EnvOrDefault("VAPID_PRIVATE_KEY", ""),
			VAPID_PUBLIC_KEY:           EnvOrDefault("VAPID_PUBLIC_KEY", ""),
			VAPID_EMAIL:                EnvOrDefault("VAPID_EMAIL", ""),
		},
		STAGING: {
			LOCAL_PORT:                 EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", ""),
			QUEUE_ACCOUNT_NAME:         EnvOrDefault("QUEUE_ACCOUNT_NAME", ""),
			QUEUE_ACCOUNT_KEY:          EnvOrDefault("QUEUE_ACCOUNT_KEY", ""),
			MONGO_DB_CONNECTION_STRING: EnvOrDefault("MONGO_DB_CONNECTION_STRING", ""),
			VAULT_URI:                  EnvOrDefault("VAULT_URI", ""),
			NOTIFY_BASE_URL:            EnvOrDefault("NOTIFY_BASE_URL", ""),
			VAPID_PRIVATE_KEY:          EnvOrDefault("VAPID_PRIVATE_KEY", ""),
			VAPID_PUBLIC_KEY:           EnvOrDefault("VAPID_PUBLIC_KEY", ""),
			VAPID_EMAIL:                EnvOrDefault("VAPID_EMAIL", ""),
		},
	}
	env_vars[PROD] = env_vars[STAGING]
	stage := Stage()
	return env_vars[stage]
}
