package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository"
	values "api/internal/domain/values/todo_values"
	models "api/internal/io_infra/database/models"
	"context"
)

type ITodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo, values.TaskId[int]]] interface {
	CreatTodo(ctx context.Context, entity_todo entity.Todo) (entity.Todo, error)
	GetAllTodo(ctx context.Context, is_sorting bool) ([]entity.Todo, error)
	DeleteTodo(ctx context.Context, target_task_id values.TaskId[int]) ([]entity.Todo, bool, error) //除去した残りを表示するためにTodoを配列で返す
	FindAll(ctx context.Context, query string, is_sorting bool) ([]entity.Todo, error)
	UpdateTodo(ctx context.Context, target_todo entity.Todo) (entity.Todo, error)
}
