package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository"
	values "api/internal/domain/values/todo_values"
	models "api/internal/io_infra/database/models"
	"context"
)

type TodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo, values.TaskId[int]]] struct {
	repository repoType
}

func NewTodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo, values.TaskId[int]]](repo repoType) *TodoUseCase[repoType] {
	todo_usecase := &TodoUseCase[repoType]{
		repository: repo,
	}
	return todo_usecase
}

func (tuc *TodoUseCase[repoType]) CreateTodo(ctx context.Context, entity_todo entity.Todo) (entity.Todo, error) {

	inserted_todo, err := tuc.repository.Create(ctx, entity_todo)
	if err != nil {
		return entity.Todo{}, err
	}
	return inserted_todo, nil
}

func (tuc *TodoUseCase[repoType]) GetAllTodo(ctx context.Context, is_sorting bool) ([]entity.Todo, error) {

	todos, err := tuc.repository.GetAll(ctx)
	if err != nil {
		return make([]entity.Todo, 0), err
	}
	if is_sorting {
		sort := NewSort(RecSort)
		sorted_todos := sort.Execute(ctx, todos)
		return sorted_todos, nil
	}
	return todos, nil
}

func (tuc *TodoUseCase[repoType]) DeleteTodo(ctx context.Context, target_task_id values.TaskId[int]) ([]entity.Todo, bool, error) { // （残りのタスク、削除成功かどうかの結果、エラー）

	result := []entity.Todo{}
	ok, err := tuc.repository.Delete(ctx, target_task_id)
	if !ok {
		return result, false, err
	}
	todos, err := tuc.repository.GetAll(ctx)
	if err != nil {
		return result, ok, err
	}
	result = append(result, todos...)
	return result, true, nil
}

func (tuc *TodoUseCase[repoType]) FindAll(ctx context.Context, query string, is_sorting bool) ([]entity.Todo, error) {

	todos, err := tuc.repository.FindAll(ctx, query)
	if err != nil {
		return make([]entity.Todo, 0), err
	}

	// sort function
	if is_sorting {
		sort := NewSort(RecSort)
		sorted_todo := sort.Execute(ctx, todos)
		return sorted_todo, nil
	}

	return todos, nil
}

func (tuc *TodoUseCase[repoType]) UpdateTodo(ctx context.Context, target_todo entity.Todo) (entity.Todo, error) {

	updated_todo_with_id, err := tuc.repository.Update(ctx, target_todo)
	if err != nil {
		return updated_todo_with_id, err
	}
	return updated_todo_with_id, nil
}
