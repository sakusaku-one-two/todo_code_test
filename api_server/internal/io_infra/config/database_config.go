package config

import (
	"api/util"
)

// Mysqlへ接続するための情報
var (
	DB_ADDR     string = util.GetEnv("", "")
	DB_PORT     string = util.GetEnv("DB_PORT", "0")
	DB_USER     string = util.GetEnv("DB_USER", "")
	DB_PASSWORD string = util.GetEnv("MYSQL_ROOT_PASSWORD", "")
)
