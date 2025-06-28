package todo_values

type TaskId[valueType comparable] struct {
	value valueType
}

// IOからのDTOの値から作成する専用
func NewTaskId[valueType int | uint | int32 | string](value valueType) (TaskId[valueType], error) {
	return TaskId[valueType]{value}, nil
}

func (t TaskId[valueType]) GetValue() valueType {
	return t.value
}
