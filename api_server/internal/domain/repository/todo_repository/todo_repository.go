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

/*
	Todoのリポジトリ
	シングルトンとして振る舞う。
*/

var todo_repository *TodoRepository // シングルトンとして扱う

type TodoRepository struct {
	driver *sql.DB
}

func NewTodoRepostory() (*TodoRepository, error) { // シングルトンを返却
	if todo_repository != nil && todo_repository.driver != nil {
		return todo_repository, nil
	}

	todo_repository = new(TodoRepository)

	db, err := driver.NewMySqlDriver()
	if err != nil {
		return nil, err
	}
	todo_repository.driver = db

	go todo_repository.self_mangement_ping() //dirverの死活監視を行うゴルーチン
	return todo_repository, nil
}

func (tr *TodoRepository) Create(insert_target entity.Todo) (entity.Todo, error) {

	todo := models.Todo{
		Title:       insert_target.Title.GetValue(),
		Description: insert_target.Description.GetValue(),
		LimitTime:   insert_target.Limit.GetValue(),
		StatusNo:    0,
	}

	ctx := context.Background()
	if err := todo.Insert(ctx, tr.driver, boil.Infer()); err != nil {
		return tr.model_to_entity(todo), err
	}
	return tr.model_to_entity(todo), nil
}

func (tr *TodoRepository) GetAll() ([]entity.Todo, error) {
	ctx := context.Background()

	var todos []*models.Todo

	if err := models.NewQuery(
		qm.From(models.TableNames.Todo),
		qm.Where("is_activate = true"), // 論理削除されてないレコードの抽出
	).Bind(ctx, tr.driver, &todos); err != nil {
		return nil, err
	}

	var entity_todos = []entity.Todo{}
	for _, todo := range todos {
		entity_todos = append(entity_todos, tr.model_to_entity(*todo))
	}

	return entity_todos, nil

}

func (tr *TodoRepository) FindAll(query string) ([]entity.Todo, error) {
	var todos []models.Todo
	ctx := context.Background()
	if err := models.NewQuery(
		qm.From(models.TableNames.Todo),
		qm.Where("title Like '%?%'", query),
	).Bind(ctx, tr.driver, &todos); err != nil {
		return nil, err
	}

	var entity_todos = []entity.Todo{}
	for _, todo := range todos {
		entity_todos = append(entity_todos, tr.model_to_entity(todo))
	}

	return entity_todos, nil
}

func (tr *TodoRepository) Update(todo entity.Todo) (entity.Todo, error)

func (tr *TodoRepository) Delete(id values.TaskId[int]) (bool, error) {

	ctx := context.Background()
	new_tx, err := tr.driver.BeginTx(ctx, nil)

	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error())
		return false, err
	}

	defer func() {
		if new_tx == nil {
			return
		}

		if err != nil {
			new_tx.Rollback()
		} else {
			new_tx.Commit()
		}
	}()

	err = models.NewQuery(
		qm.From(models.TableNames.Todo),
		qm.Where("id = ?", id.GetValue()),
	).Bind(ctx, new_tx, nil)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (tr *TodoRepository) model_to_entity(todo models.Todo) entity.Todo {

	Id, _ := values.NewTaskId(todo.ID)
	Title, _ := values.NewTitle(todo.Title)
	Description, _ := values.NewDescription(todo.Description)
	Limit, _ := values.NewLimit(todo.LimitTime)
	Status, _ := values.GetTodoStatus(todo.StatusNo)

	var Is_activate bool
	if todo.IsActivate == 1 { //真偽値はDBから取得した際０または１なので条件分岐でTrue Falseに変換
		Is_activate = true
	} else {
		Is_activate = false
	}

	Update_at := todo.UpdateAt.Time
	Created_at := todo.CreatedAt

	return entity.Todo{
		Id:          Id,
		Title:       Title,
		Description: Description,
		Limit:       Limit,
		Status:      Status,
		Is_activate: Is_activate,
		Update_at:   Update_at,
		Created_at:  Created_at,
	}
}

// --------------------------------[DBとの死活監視を実施]-------------------------------//

func (tr *TodoRepository) self_mangement_ping() {
	ticker := time.NewTicker(60 * time.Second)
	for {
		// 1分おきに死活監視を実行
		<-ticker.C

		if err := tr.driver.Ping(); err != nil {
			slog.Log(context.Background(), slog.LevelInfo, err.Error())
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
}
