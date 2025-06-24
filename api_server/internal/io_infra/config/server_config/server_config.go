package config

import (
	"api/util"
	"fmt"
)

var (
	SERVER_PORT = util.GetEnv("SERVER_PORT", "8080")
	SERVER_HOST = util.GetEnv("SERVER_HOST", "localhost")
)

func GetServerAddressAndPort() string {
	return fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT)
}
