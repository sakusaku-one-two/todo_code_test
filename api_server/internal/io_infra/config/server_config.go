package config

import (
	"api/util"
)

var (
	SERVER_PORT = util.GetEnv("SERVER_PORT", "8080")
	SERVER_ADDR = util.GetEnv("SERVER_ADDR", "localhost")
)
