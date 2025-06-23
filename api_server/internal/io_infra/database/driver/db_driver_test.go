package driver

import (
	m "api/internal/io_infra/database/models"
	"context"
	"fmt"
	"testing"
)

func Test_connection_test(t *testing.T) {
	conn, err := NewMySqlDriver()

	defer func() {
		if conn != nil {
			conn.Close()
			fmt.Println("mysql closed!!")
		}
	}()

	if err != nil {
		t.Log("DB 接続失敗", err.Error())
		return
	}

	if err := conn.Ping(); err != nil {
		t.Log(
			"DBとの接続に問題あり", err.Error(),
		)
		return
	}

	ctx := context.Background()

	todo, err := m.FindTodo(ctx, conn, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(todo)

}
