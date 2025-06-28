package todo_values

import (
	"fmt"
	"testing"
	"time"
)

func Test_limit(t *testing.T) {
	fmt.Println("Limitのテストを開始")
	defer fmt.Println("Limitのテスト終了")
	// 期限のテスト
	now_ := time.Now()
	limit, err := NewLimit(now_)
	fmt.Println(limit, err)
	//同一性の確認
	identify := limit == limit
	fmt.Println(fmt.Printf("同一value objectの比較 %t", identify))
	now_ = now_.Add(10 * time.Hour)
	new_limit, _ := NewLimit(now_)
	identify = limit == new_limit
	fmt.Println(fmt.Printf("別のvalue objectの比較 %t", identify))
}

func Test_Status(t *testing.T) {
	fmt.Println("Statusの検証開始")
	defer fmt.Println("Statusの検証終了")
	status, err := GetTodoStatus(COMPLETE)
	fmt.Println("COMPLETE =>", status, err)
	_, err = GetTodoStatus(10)
	if err != nil {
		fmt.Println("Statusエラー生成成功", err)
	} else {
		t.Error("Statusのエラー生成失敗")
	}

}

func Test_value(t *testing.T) {
	new_val, err := NewTaskId[int](1)
	if err != nil {
		t.Fatal()
		return
	}

	fmt.Println(new_val)

}
