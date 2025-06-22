package todo_values

import (
	"github.com/google/uuid"
)

type TaskId struct {
	value string
}

// IOからのDTOの値から作成する専用
func NewTaskId(value string) (TaskId, error) {
	task_id := TaskId{}
	uuid, err := uuid.Parse(value)
	if err != nil {
		return task_id, err
	}
	task_id.value = uuid.String()
	return task_id, nil
}

// 新規作成
func GenerateTaskId() TaskId {
	return TaskId{
		value: uuid.New().String(),
	}
}
