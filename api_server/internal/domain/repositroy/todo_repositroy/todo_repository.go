package todo_repository

import (
	db "api/internal/io_infra/database/driver"
	values "api/internal/values/todo_values"
)

type TodoRepository struct {
	connectionHandler db.ConnectionHandler[db.DriverTypeAsMySql]
}

func (tr *TodoRepository) Create(insert_target values.Todo) (bool, error) {

	return true, nil
}
