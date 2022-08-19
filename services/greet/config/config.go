package config

import (
	configlib "github.com/eugbyte/monorepo/libs/config"
)

type vars struct {
	LOCAL_PORT string
}

// redeclare variables to avoid confusion between current config pkg and configLib pkg
type STAGE = configlib.STAGE

var DEV = configlib.DEV
var STAGING = configlib.STAGING
var PROD = configlib.PROD
var Stage func() configlib.STAGE = configlib.Stage
var EnvOrDefault func(key string, defaultValue string) string = configlib.EnvOrDefault

func New() vars {
	var env_vars map[STAGE]vars = map[STAGE]vars{
		DEV: {
			LOCAL_PORT: EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
		},
		STAGING: {
			LOCAL_PORT: EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", ""),
		},
	}
	env_vars[PROD] = env_vars[STAGING]
	stage := Stage()
	return env_vars[stage]
}
