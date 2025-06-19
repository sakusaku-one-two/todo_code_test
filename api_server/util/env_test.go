package util

import (
	"testing"
)

const (
	ENV_TEST_NAME     = "ENV_TEST_NAME"
	ENV_TEST_VALUE    = "ENV_TEST_VALUE"
	ENV_DEFAULT_VALUE = "ENV_DEFULET_VALUE"
)

func Test_env(t *testing.T) {
	// 環境変数の取得テスト
	defer func() {
		// 実行後の環境変数のクリーンアップ
		if DeleteEnv(ENV_TEST_NAME) {
			t.Log("環境変数の削除完了")
		} else {
			t.Log("環境変数の削除失敗")
		}
	}()

	// 環境変数の設定テスト
	if ok := SetEnv(ENV_TEST_NAME, ENV_TEST_VALUE); ok {
		t.Log("環境変数の設定完了")
	} else {
		t.Log("環境変数の設定失敗")
	}

	//　環境変数の取得テスト
	env_result := GetEnv(ENV_TEST_NAME, ENV_DEFAULT_VALUE)

	switch env_result {
	case ENV_DEFAULT_VALUE:
		t.Log("環境変数の取得に失敗")
	case ENV_TEST_VALUE:
		t.Log("環境変数の取得に成功")
	default:
		t.Log("???")
	}

}
