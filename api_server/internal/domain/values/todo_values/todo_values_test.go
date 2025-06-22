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
	sample_id := GenerateTaskId()
	sample_copy_id, _ := NewTaskId(sample_id.value) // Ioからの値を想定
	//同一性の確認1
	if sample_id == sample_copy_id {
		// 比較演算子で問題なく同一性を保証できるかの確認
		fmt.Println("sample == sample_copy_id OK!!")
	} else {
		t.Error("TaskIDの比較演算子（==）を用いた比較が失敗")
	}

	other_id := GenerateTaskId()

	// 同一性の確認2
	if sample_id != other_id {
		fmt.Println("sample_id != other_id OK!!")
	} else {
		t.Error("TasiKIdの比較演算(!=)を用いた比較が失敗")
	}

	fmt.Println("Ioから無効な値が入った際のエラー確認開始")
	const fake_task_id_value = "onaka hetta!!"

	_, err := NewTaskId(fake_task_id_value)

	if err != nil {
		fmt.Println("NewTaskIdのエラー生成成功", err)
	} else {
		t.Error("UUID型として変換出来ない文字列をUUIDとして認識してしまった")
	}

	fmt.Println("Ioから無効な値が入った際のエラー確認完了")

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
