package todo_repository

import (
	entity "api/internal/domain/entitys/todo_entity"
	values "api/internal/domain/values/todo_values"
	db_modle "api/internal/io_infra/database/models"
)

type IRepository[modelT db_modle.Todo, domainT entity.Todo] interface {
	Create(domainT) (domainT, bool)
	Update(domainT) (domainT, bool)
	GetAll() ([]domainT, error)
	FindAll(string) ([]domainT, error)
	Delete(values.TaskId) bool
}
