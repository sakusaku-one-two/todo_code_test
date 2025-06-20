package database_driver

import (
	"testing"
)

func Test_connection_test(t *testing.T) {
	conn, err := NewMySqlConnectionHandler()
	if err != nil {
		t.Log("DB 接続失敗", err.Error())
	}

	if ok := conn.Ping(); ok {
		t.Log(
			"DBとの接続に問題なし",
		)
	}

	defer conn.Close()

}
