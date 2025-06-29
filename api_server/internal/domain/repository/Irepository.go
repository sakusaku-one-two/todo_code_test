package repository

import (
	values "api/internal/domain/values"
	"context"
)

type IRepository[modelT any, domainT any, id values.IValue[int]] interface {
	Create(context.Context, domainT) (domainT, error)
	Update(context.Context, domainT) (domainT, error)
	GetAll(context.Context) ([]domainT, error)
	FindAll(ctx context.Context, query string) ([]domainT, error)
	Delete(context.Context, id) (bool, error)
}
