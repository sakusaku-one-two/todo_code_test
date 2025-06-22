package todo_values

import (
	"github.com/google/uuid"
)

type TaskId struct {
	value uuid.UUID
}

func NewTaskId() (TaskId, error) {
	return TaskId{
		value: uuid.New(),
	}, nil
}
