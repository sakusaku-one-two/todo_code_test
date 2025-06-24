package todo_repository

import (
	entity "api/internal/domain/entitys/todo_entity"
	values "api/internal/domain/values/todo_values"
	driver "api/internal/io_infra/database/driver"
	models "api/internal/io_infra/database/models"
	"context"
	"database/sql"
	"log/slog"
	"sync"
	"time"
)

type TodoRepository struct {
	driver *sql.DB
}

func NewTodoRepostory() (TodoRepository, error) {
	var todo_repo TodoRepository
	db, err := driver.NewMySqlDriver()
	if err != nil {
		return todo_repo, err
	}
	todo_repo.driver = db
	// dirverの死活監視を行うゴルーチン
	go todo_repo.self_mangement_ping()
	return todo_repo, nil
}

func (tr *TodoRepository) Create(insert_target entity.Todo) (bool, error) {
	

	todo := models.Todo{
		Title:       insert_target.Title.GetValue(),
		Description: insert_target.Description.GetValue(),
		Limit: 
	}

	ctx := context.Background()
	todo.Insert(ctx)

	return true, nil
}

func (tr *TodoRepository) Read(id values.TaskId) (entity.Todo, error) {

}

func (tr *TodoRepository) self_mangement_ping() {
	ticker := time.NewTicker(60 * time.Second)
	for {
		// 1分おきに死活監視を実行
		helth_check_time := <-ticker.C

		if err := tr.driver.Ping(); err != nil {
			slog.Log(context.Background(), slog.LevelWarn, err.Error(), helth_check_time)
			tr.driver_update() //pingでダメだったら新しいドライバーコネクションに交換する
		}

	}
}

func (tr *TodoRepository) driver_update() {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	old_db_driver := tr.driver // 既存のコネクションを一時的に構造体から退避

	new_db_driver, err := driver.NewMySqlDriver()
	if err != nil {
		return
	}
	tr.driver = new_db_driver //新しいコネクションに交換
	old_db_driver.Close()
	return
}
