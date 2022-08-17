package config

import (
	configlib "github.com/eugbyte/monorepo/libs/config"
)

type vars struct {
	LOCAL_PORT                 string
	QUEUE_ACCOUNT_NAME         string
	QUEUE_ACCOUNT_KEY          string
	MONGO_DB_CONNECTION_STRING string
	VAPID_PUBLIC_KEY           string
	VAPID_SENDER_EMAIL         string
	VAULT_URI                  string
	NOTIFY_PRODUCER_URL        string
}

// redeclare variables to avoid confusion between current config pkg and configLib pkg
type STAGE = configlib.STAGE

var DEV = configlib.DEV
var STAGING = configlib.STAGING
var PROD = configlib.PROD
var Stage func() configlib.STAGE = configlib.Stage
var EnvOrDefault func(key string, defaultValue string) string = configlib.EnvOrDefault
var QueueBaseURL func(stage STAGE, accountName string) string = configlib.QueueBaseURL

var ENV_VARS = map[STAGE]vars{
	DEV: {
		LOCAL_PORT:                 EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
		QUEUE_ACCOUNT_NAME:         "devstoreaccount1",
		QUEUE_ACCOUNT_KEY:          "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==",
		MONGO_DB_CONNECTION_STRING: `mongodb://localhost:C2y6yDjf5%2FR%2Bob0N8A7Cgv30VRDJIWEHLM%2B4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw%2FJw%3D%3D@localhost:10255/admin?ssl=true`,
		VAPID_PUBLIC_KEY:           "BPlL5OTZwtW-0-4pQXmobTgX6URszc9-UKoTTvpvInhUlPHorlDM8y04J-rrErlQXMVH7_Us983mNmmwsb-z53U",
		VAPID_SENDER_EMAIL:         "eugenetham1994@gmail.com",
		VAULT_URI:                  "https://kv-notify-secrets-stg.vault.azure.net",
		NOTIFY_PRODUCER_URL:        "http://localhost:7071/api/notifications",
	},
	STAGING: {
		LOCAL_PORT:                 EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", ""),
		QUEUE_ACCOUNT_NAME:         EnvOrDefault("QUEUE_ACCOUNT_NAME", ""),
		QUEUE_ACCOUNT_KEY:          EnvOrDefault("QUEUE_ACCOUNT_KEY", ""),
		MONGO_DB_CONNECTION_STRING: EnvOrDefault("MONGO_DB_CONNECTION_STRING", ""),
		VAPID_PUBLIC_KEY:           "BPlL5OTZwtW-0-4pQXmobTgX6URszc9-UKoTTvpvInhUlPHorlDM8y04J-rrErlQXMVH7_Us983mNmmwsb-z53U",
		VAPID_SENDER_EMAIL:         "eugenetham1994@gmail.com",
		VAULT_URI:                  EnvOrDefault("VAULT_URI", ""),
		NOTIFY_PRODUCER_URL:        "http://localhost:7071/api/notifications",
	},
}
