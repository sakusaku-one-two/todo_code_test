package my_sql_config

import (
	"api/util"
)

// Mysqlへ接続するための情報 環境変数が存在しなければ、第二引数のデフォルト値が代入される
var (
	DB_ADDR          string = util.GetEnv("MYSQL_HOSTNAME", "localhost")
	DB_PORT          string = util.GetEnv("MYSQL_PORT", "3306")
	DB_USER          string = util.GetEnv("MYSQL_USER", "mysql_user")
	DB_ROOT_PASSWORD string = util.GetEnv("MYSQL_ROOT_PASSWORD", "root_password")
	DB_PASSWORD      string = util.GetEnv("MYSQL_PASSWORD", "todo_database_password")
	DB_NAME          string = util.GetEnv("MYSQL_DATABASE", "todo_database")
	DB_MAX_CONN      int    = util.GetEnvAsInt("MYSQL_MAX_CONN", 5)
)
