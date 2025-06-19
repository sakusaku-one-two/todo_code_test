package config

import (
	"api/util"
	"os"
)

// Mysqlへ接続するための情報
var (
	DB_ADDR     string = util.GetEnv("DB_ADDR", "")
	DB_PORT     string = util.GetEnv("DB_PORT", "0")
	DB_USER     string = util.GetEnv("DB_USER", "")
	DB_PASSWORD string = util.GetEnv("DB_PASSWORD", "")
)

func init() {

	cuurent_dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	util.SetEnv("XDG_CONFIG_HOME", cuurent_dir)
}
