package type_entitys

import (
	todo "api/internal/domain/values/todo_values"
	"time"
)

type Todo struct {
	Id          todo.TaskId[int]
	Title       todo.Title
	Description todo.Description
	Limit       todo.Limit
	Status      todo.Status
	Is_activate bool
	Update_at   time.Time
	Created_at  time.Time
}
