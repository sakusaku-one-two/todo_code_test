package todo_values

import (
	"fmt"
	"time"
)

type Limit struct {
	value time.Time
}

func NewLimit(limit_time time.Time) (limit Limit, err error) {
	if limit_time.Before(time.Now()) { // 現在時刻より前の時刻の場合はエラー
		return limit, fmt.Errorf("the deadline rule is that it must be later than now: %s >= %s", limit_time, time.Now())
	}
	limit.value = limit_time
	return limit, err
}

func (l *Limit) GetValue() time.Time {
	return l.value
}
