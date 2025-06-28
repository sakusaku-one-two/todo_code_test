package todo_values

type TaskId[valueType comparable] struct {
	value valueType
}

// IOからのDTOの値から作成する専用
func NewTaskId(value int) (TaskId[int], error) {
	return TaskId[int]{value}, nil
}

func (t TaskId[valueType]) GetValue() valueType {
	return t.value
}
