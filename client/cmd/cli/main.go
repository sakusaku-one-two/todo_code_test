package main

import (
	v1 "api/internal/grpc_gen/todo/v1"
	todov1connect "api/internal/grpc_gen/todo/v1/todov1connect"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var input_value string

func main() {

	TodoServiceClient := todov1connect.NewTodoServiceClient(
		&http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
					// If you're also using this client for non-h2c traffic, you may want to
					// delegate to tls.Dial if the network isn't TCP or the addr isn't in an
					// allowlist.
					return net.Dial(network, addr)
				},
			},
		},
		"http://localhost:8080",
		connect.WithGRPC(),
	)

	for {

		fmt.Println("'削除'・'作成'・'更新'(完了)・'検索'・'取得'・'終了'のどれかを入力してください。 ctl+cで強制終了")
		fmt.Scan(&input_value)

		switch input_value {
		case "作成":
			CreateTodo(TodoServiceClient)
		case "取得":
			GetAll(TodoServiceClient)
		case "検索":
			FindSerch(TodoServiceClient)
		case "更新":
			Update(TodoServiceClient)
		case "削除":
			Delete(TodoServiceClient)
		case "終了":
			return
		}

	}
}

//----------------------[以下ロジック]--------------------//

func PrintTodo(idx int, todo *v1.Todo) {
	idx += 1

	fmt.Println(idx, "___________________________________________________")
	fmt.Println("タイトル", todo.Title)
	fmt.Println("説明　　", todo.Description)
	fmt.Println("期限　　", ToDate(todo.LimitTime))
	fmt.Println("状態　　", GetStatus(todo.Status))
	fmt.Println("作成日　", ToDate(todo.GetCreatedAt()))
	fmt.Println("________________________________________________________")
}

func ToDate(target *timestamppb.Timestamp) string {
	target_time := target.AsTime()
	return fmt.Sprintf("%d年%d月%d日", target_time.Year(), target_time.Month(), target_time.Day())
}

func ToInt(target string) int {

	resutl, err := strconv.Atoi(target)
	if err != nil {
		var temp string
		for {
			fmt.Println("数値で再度入力してください。")
			fmt.Scan(&temp)
			scaned, err := strconv.Atoi(temp)
			if err == nil {
				return scaned
			}
		}
	}
	return resutl
}

func CreateTodo(todo_service_client todov1connect.TodoServiceClient) {
	var title, descrpition string
	var limit_month, limit_day string
	var limit_mouth_as_int, limit_day_as_int int

	fmt.Println("タイトルを入力してください。")
	fmt.Scan(&title)
	fmt.Println("説明を入力してください。")
	fmt.Scan(&descrpition)
	fmt.Println("期限の月(1~12)を半角数値で指定してください。")
	fmt.Scan(&limit_month)
	limit_mouth_as_int = ToInt(limit_month)
	fmt.Println("期限の日を半角数値で入力してください。")
	fmt.Scan(&limit_day)
	limit_day_as_int = ToInt(limit_day) + 1
	year := time.Now().Local().Year()
	limit_date := time.Date(year, time.Month(limit_mouth_as_int), limit_day_as_int, 0, 0, 0, 0, time.Local)

	if limit_date.Before(time.Now()) {
		fmt.Println("期限の指定は今日以降の日付を入れてください。")
		return
	}

	res, err := todo_service_client.CreateTodo(context.Background(),
		connect.NewRequest(&v1.CreateTodoRequest{RequestTodo: &v1.Todo{
			Title:       title,
			Description: descrpition,
			LimitTime:   timestamppb.New(limit_date),
			Status:      v1.Status_INCOMPLETE,
		}}),
	)
	if err != nil {
		fmt.Println("サーバーエラーのため作成できませんでした")
		return
	}

	PrintTodo(1, res.Msg.CreatedTodo)

}

func GetAll(todo_service_client todov1connect.TodoServiceClient) {
	req, err := todo_service_client.GetAllTodo(
		context.Background(),
		connect.NewRequest(&v1.GetALLRequest{Request: "", IsSort: true}),
	)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for idx, todo := range req.Msg.Result {
		PrintTodo(idx, todo)
	}
}

func GetStatus(status v1.Status) string {
	if status == 1 {
		return "完了"
	}
	return "未完了"
}

func FindSerch(todo_service_client todov1connect.TodoServiceClient) {

	conn := todo_service_client.FindTodo(context.Background())

	defer func() {
		conn.CloseRequest()
	}()
	var input_var string
	for {

		fmt.Println("検索するタイトルを入力してください。（部分検索可能）")

		fmt.Scan(&input_var)
		conn.Send(&v1.SearchRequest{Query: input_var})

		res, err := conn.Receive()
		if err != nil {
			fmt.Println("err!")
			return
		}

		for idx, todo := range res.GetResult() {
			PrintTodo(idx, todo)
		}
		fmt.Println("検索終了しますか？ 'yes'で終了")
		fmt.Scan(&input_var)
		if input_var == "yes" {
			return
		}

	}
}

func Update(todo_server_client todov1connect.TodoServiceClient) {
	ctx := context.Background()
	res, err := todo_server_client.GetAllTodo(ctx, connect.NewRequest(&v1.GetALLRequest{Request: "", IsSort: true}))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	todos := res.Msg.GetResult()

	if len(todos) == 0 {
		fmt.Println("現在登録されたTODOはありません。終了します。")
		return
	}

	for idx, todo := range todos {
		PrintTodo(idx, todo)
	}

	fmt.Println("右上のindex番号を指定してください。")
	var input_string string
	fmt.Scan(&input_string)
	idx_no := ToInt(input_string) - 1

	if idx_no < 0 {
		fmt.Println("indexの範囲指定外です。")
		return
	}

	if len(todos) < idx_no {
		fmt.Println("indexの範囲指定外です。")
		return
	}

	target_todo := todos[idx_no]

	target_todo.Status = v1.Status_COMPLETE

	ctx = context.Background()
	update_response, err := todo_server_client.UpdateTodo(ctx, connect.NewRequest(&v1.UpdateTodoRequest{Todo: target_todo}))

	if err != nil {
		fmt.Println("更新に失敗しました。", update_response, err.Error())
		return
	}

	fmt.Println("タスクを更新しました。")
}

func Delete(todo_service_client todov1connect.TodoServiceClient) {
	ctx := context.Background()

	req, err := todo_service_client.GetAllTodo(
		context.Background(),
		connect.NewRequest(&v1.GetALLRequest{Request: "", IsSort: true}),
	)

	if err != nil {
		fmt.Println("タスクの取得に失敗しました。")
		return
	}

	todos := req.Msg.GetResult()

	if len(todos) == 0 {
		fmt.Println("現在登録されたTODOはありません。終了します。")
		return
	}

	for idx, todo := range todos {
		PrintTodo(idx, todo)
	}

	fmt.Println("削除する対象のTodo右上のindex番号を指定してください。")
	var input_string string
	fmt.Scan(&input_string)
	idx_no := ToInt(input_string) - 1

	if idx_no < 0 {
		fmt.Println("indexの範囲指定外です。")
		return
	}

	if len(todos) < idx_no {
		fmt.Println("indexの範囲指定外です。")
		return
	}

	target_todo := todos[idx_no]

	res, err := todo_service_client.DeleteTodo(ctx, connect.NewRequest(&v1.DeleteTodoRequest{Id: *target_todo.Id}))
	if err != nil || !res.Msg.Result {
		fmt.Println("削除に失敗しました。")
		return
	}

	fmt.Println("削除に成功しました")

}
