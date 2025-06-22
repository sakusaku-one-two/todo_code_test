package todo_values

import (
	"github.com/google/uuid"
)

type TaskId struct {
	value string
}

func NewTaskId() (TaskId, error) {
	return TaskId{
		value: uuid.New().String(),
	}, nil
}
