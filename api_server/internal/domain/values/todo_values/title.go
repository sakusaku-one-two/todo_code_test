package todo_values

import (
	"fmt"
)

const (
	// 1文字以上５０文字以下
	RULE_OVER_COUNT  = 1
	RULE_UNDER_COUNT = 50
)

type Title struct {
	value string
}

func NewTitle(value string) (Title, error) {

	if len(value) >= RULE_OVER_COUNT && len(value) <= RULE_UNDER_COUNT {
		return Title{}, fmt.Errorf("must be length over %d and under %d :: target value => %s , length => %d", RULE_OVER_COUNT, RULE_UNDER_COUNT, value, len(value))
	}

	return Title{
		value,
	}, nil
}
