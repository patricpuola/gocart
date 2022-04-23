package config

import (
	"strconv"
)

const NOT_VERBOSE = 0
const VERBOSE = 1
const VERY_VERBOSE = 2

var Config map[string]string
var Verbosity int = NOT_VERBOSE

func GetInt(key string) int {
	value := Config[key]
	integer, _ := strconv.Atoi(value)
	return integer
}

func GetString(key string) string {
	return Config[key]
}

func IsTest() bool {
	return Config["env"] == "test"
}

func IsProd() bool {
	return Config["env"] == "prod"
}

func IsVerbose() bool {
	return Verbosity >= VERBOSE
}

func IsVeryVerbose() bool {
	return Verbosity >= VERY_VERBOSE
}
