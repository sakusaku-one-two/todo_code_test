package repository

import (
	values "api/internal/domain/values"
)

type IRepository[modelT any, domainT any, id values.IValue[int]] interface {
	Create(domainT) (domainT, error)
	Update(domainT) (domainT, error)
	GetAll() ([]domainT, error)
	FindAll(query string) ([]domainT, error)
	Delete(id) (bool, error)
}
