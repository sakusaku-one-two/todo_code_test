package todo_values

import (
	"fmt"
)

const (
	INCOMPETE = 0
	COMPLETE  = 1
)

var (
	statsuMap        = map[int]Status{}
	reverseMap       = map[Status]int{}
	MAX_STATUS_COUNT int
)

func init() {
	statsuMap[COMPLETE] = Status{value: "COMPLETE"}
	statsuMap[INCOMPETE] = Status{value: "INCOMPLETE"}
	MAX_STATUS_COUNT = len(statsuMap) - 1
	for key, value := range statsuMap {
		reverseMap[value] = key
	}
}

type Status struct {
	value string
}

func GetTodoStatus(enum_constraint_no int) (Status, error) {
	if -1 > enum_constraint_no || enum_constraint_no > MAX_STATUS_COUNT {
		return Status{}, fmt.Errorf("out of enum range: provide key is %d", enum_constraint_no)
	}
	//読み込み専用なので、データ競合を防ぐLockはかけない。
	return statsuMap[enum_constraint_no], nil
}

func GetTodoStatusNumber(status Status) int {
	return reverseMap[status]
}
