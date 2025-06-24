package driver

import (
	m "api/internal/io_infra/database/models"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	qm "github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	// 初期化スクリプトでinsertした値の取得
	todo, err := m.FindTodo(ctx, conn, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(todo)

	var todos []*m.Todo

	m.NewQuery(
		qm.From(m.TableNames.Todo),
	).Bind(ctx, conn, &todos)

	for _, todo := range todos {
		fmt.Println(*todo)
	}

	new_todo := &m.Todo{
		Title:       "sample_title_99",
		Description: "sample_description_99",
		LimitTime:   time.Now(),
	}

	err = new_todo.Insert(ctx, conn, boil.Infer())
	if err != nil {
		fmt.Println(err.Error())
	}

}
