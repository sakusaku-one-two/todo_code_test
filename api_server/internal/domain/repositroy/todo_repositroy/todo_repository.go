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

	"github.com/volatiletech/sqlboiler/v4/boil"
	qm "github.com/volatiletech/sqlboiler/v4/queries/qm"
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
		LimitTime:   insert_target.Limit.GetValue(),
	}
	ctx := context.Background()
	if err := todo.Insert(ctx, tr.driver, boil.Infer()); err != nil {
		return false, err
	}

	return true, nil
}

func (tr *TodoRepository) Read(id values.TaskId) (entity.Todo, error) {

}

func (tr *TodoRepository) GetAll() ([]*entity.Todo, error) {}

func (tr *TodoRepository) FindAll(query string) ([]entity.Todo, error) {
	var todos []entity.Todo
	ctx := context.Background()
	if err := models.NewQuery(
		qm.From(models.TableNames.Todo),
		qm.Where("title Like %?%", query),
	).Bind(ctx, tr.driver, todos); err != nil {
		return todos, err
	}

	return todos, nil

}

func (tr *TodoRepository) self_mangement_ping() {
	ticker := time.NewTicker(60 * time.Second)
	for {
		// 1分おきに死活監視を実行
		helth_check_time := <-ticker.C

		if err := tr.driver.Ping(); err != nil {
			slog.Log(context.Background(), slog.LevelWarn, err.Error(), helth_check_time)
			tr.driver_recreate() //pingでダメだったら新しいドライバーコネクションに交換する
		}

	}
}

func (tr *TodoRepository) driver_recreate() {
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
