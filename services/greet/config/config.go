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

var ENV_VARS = map[STAGE]vars{
	DEV: {
		LOCAL_PORT: configlib.EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
	},
	STAGING: {
		LOCAL_PORT: configlib.EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
	},
	PROD: {
		LOCAL_PORT: configlib.EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", "8080"),
	},
}
