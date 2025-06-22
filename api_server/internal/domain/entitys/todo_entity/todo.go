package type_entitys

import (
	todo "api/internal/domain/values/todo_values"
	"time"
)

type Todo struct {
	id          todo.TaskId
	title       todo.Title
	description todo.Description
	limit       todo.Limit
	status      todo.Status
	is_activate bool
	update_at   time.Time
	created_at  time.Time
}
