package todo_values

type TaskId struct {
	value int
}

// IOからのDTOの値から作成する専用
func NewTaskId(value int) (TaskId, error) {
	return TaskId{value}, nil
}

// 新規作成
