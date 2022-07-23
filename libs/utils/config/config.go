package config

import (
	"fmt"
	"log"
	"os"
)

type STAGE int

const (
	DEV STAGE = iota
	STAGING
	PROD
)

func EnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Stage() STAGE {
	var stage = EnvOrDefault("STAGE", "dev")
	stageMap := map[string]STAGE{
		"dev":  DEV,
		"stg":  STAGING,
		"prod": PROD,
	}
	if stageMap == nil {
		log.Panicln("STAGE environment variable must be ['dev', 'stg', 'prod']")
	}
	return stageMap[stage]
}

type vars struct {
	LOCAL_PORT                 string
	QUEUE_ACCOUNT_NAME         string
	QUEUE_ACCOUNT_KEY          string
	MONGO_DB_CONNECTION_STRING string
}

var ENV_VARS = map[STAGE]vars{
	DEV: {
		LOCAL_PORT:                 EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
		QUEUE_ACCOUNT_NAME:         "devstoreaccount1",
		QUEUE_ACCOUNT_KEY:          "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==",
		MONGO_DB_CONNECTION_STRING: `mongodb://localhost:C2y6yDjf5%2FR%2Bob0N8A7Cgv30VRDJIWEHLM%2B4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw%2FJw%3D%3D@localhost:10255/admin?ssl=true`,
	},
	STAGING: {},
	PROD:    {},
}

func QueueBaseURL(stage STAGE, accountName string) string {
	if stage == DEV {
		return fmt.Sprintf("%s/%s", "http://127.0.0.1:10001", accountName)
	} else {
		return fmt.Sprintf("https://%s.queue.core.windows.net", accountName)
	}
}
