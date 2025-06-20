package config

import (
	"errors"
	"testing"
)

func Test_database_config(t *testing.T) {
	if DB_ADDR != "" &&
		DB_PASSWORD != "" &&
		DB_PORT != "" {
		t.Log("database env ok!")

	} else {
		t.Error(errors.New("データベースの環境変数の取得に失敗"))
	}
}
