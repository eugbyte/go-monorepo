package config

import (
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
	return stageMap[stage]
}
