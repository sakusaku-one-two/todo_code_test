// package util

// // ---------------------------------------------[コールバック関数]-----------------------------------------------------------------//

// // この型の関数をWorkFlowのnextメソッド　に渡す
// type CmdFunction[FirstStateValue, NextStateValue any] = func(state State[FirstStateValue]) State[NextStateValue]

// // --------------------------------------------------------------------------------------------------------------//

// type WorkFlow[ValueType any] struct {
// 	state State[ValueType] // ValueTypeはState内にあるプロパティ　である　valueの型情報
// }

// func NewWorkFlow[ValueType any](first_state_value ValueType) WorkFlow[ValueType] {
// 	return WorkFlow[ValueType]{
// 		state: NewState(first_state_value),
// 	}
// }

// func (wf WorkFlow[ValueType]) ErrorState(err_state State[ValueType]) WorkFlow[ValueType] {
// 	wf.state = err_state
// 	return wf
// }

// func (wf WorkFlow[ValueType]) Next(call_back CmdFunction[ValueType, any]) WorkFlow[ValueType] {
// 	if wf.state.IsErr() {
// 		return wf
// 	}

// 	newState := call_back(wf.state)

// 	return NewWorkFlow[ValueType](newState)

// }
