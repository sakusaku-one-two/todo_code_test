package todo_repository

import (
	values "api/internal/values/todo_values"
)

type TodoRepository struct {
	driver driver.
}

func (tr *TodoRepository) Create(insert_target values.Todo) (bool, error) {

	return true, nil
}
