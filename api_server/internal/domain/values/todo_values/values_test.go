package todo_values

import (
	"fmt"
	"testing"
	"time"
)

func Test_limit(t *testing.T) {
	now_ := time.Now()
	limit, err := NewLimit(now_)
	fmt.Println(limit, err)

	//同一かの比較

	identify := limit == limit

	fmt.Println(fmt.Printf("同一value objectの比較 %t", identify))

	now_ = now_.Add(10 * time.Hour)
	new_limit, _ := NewLimit(now_)
	identify = limit == new_limit
	fmt.Println(fmt.Printf("別のvalue objectの比較 %t", identify))

}
