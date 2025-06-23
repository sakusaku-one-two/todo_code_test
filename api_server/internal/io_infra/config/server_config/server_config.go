package config

import (
	"api/util"
)

var (
	SERVER_PORT = util.GetEnv("SERVER_PORT", "8080")
	SERVER_HOST = util.GetEnv("SERVER_HOST", "localhost")
)
