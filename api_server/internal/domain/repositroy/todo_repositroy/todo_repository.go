package todo_repository

import (
	db "api/internal/io_infra/database/driver"
)

type TodoRepository struct {
	connectionHandler db.ConnectionHandler[db.DriverTypeAsMySql]
}
