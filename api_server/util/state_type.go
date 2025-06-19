package util

/*
	レール指向プログラミングの状態型の実装
*/

type State[V any] struct {
	value V
	err   error
}

// 外部コンストラクター
func NewState[V any](target_value V) State[V] {
	return State[V]{
		value: target_value,
		err:   nil,
	}
}

// エラー専用の状態を作成するコンストラクター
func NewErrorState(err error) State[any] {
	return State[any]{
		err: err,
	}
}

func (r *State[V]) GetValue() V {
	return r.value
}

func (r *State[V]) IsErr() bool {
	return r.err != nil
}

func (r *State[V]) ErrorMessage() string {
	return r.err.Error()
}
