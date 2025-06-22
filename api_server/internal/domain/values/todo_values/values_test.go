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

func Test_id(t *testing.T) {
	fmt.Println("TaskIdのテストを開始")
	defer fmt.Println("TaskIdのテスト終了")
	// IDのテスト
	sample_id, _ := GenerateTaskId()
	//同一性の確認1
	if sample_id == sample_id {
		// 比較演算子で問題なく同一性を保証できるかの確認
		fmt.Println("sample == sample_id OK!!")
	} else {
		t.Error("TaskIDの比較演算子（==）を用いた比較が失敗")
	}

	other_id, _ := GenerateTaskId()

	// 同一性の確認2
	if sample_id != other_id {
		fmt.Println("sample_id != other_id OK!!")
	} else {
		t.Error("TasiKIdの比較演算(!=)ヲ用いた比較が失敗")
	}

}
