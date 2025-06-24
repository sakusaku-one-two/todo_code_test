package util

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

/*
	環境変数に関する操作を行う
	GetEnv    ==> 環境変数を取得する
	SetEnv    ==>　環境変数を設定する
	DeleteEnv ==> 環境変数の削除
*/

func GetEnv(env_name, defult_value string) string {
	if env_data, ok := os.LookupEnv(env_name); ok {
		return env_data
	}
	return defult_value
}

func SetEnv(env_name, setup_value string) bool {
	err := os.Setenv(env_name, setup_value)
	if err != nil {
		slog.Info(
			fmt.Sprintf("環境変数 : %s の設定が失敗", env_name),
		)
		return false
	}
	return true
}

func DeleteEnv(env_name string) bool {
	if err := os.Unsetenv(env_name); err != nil {
		return false
	}
	return true
}

func GetEnvAsInt(env_name string, default_value int) int {
	env_value := GetEnv(env_name, "nothing")
	if env_value == "nothing" {
		return default_value
	}
	result, err := strconv.Atoi(env_value)
	if err != nil {
		ctx := context.Background()
		slog.Log(ctx, slog.LevelError, err.Error())
		return default_value
	}

	return result

}
