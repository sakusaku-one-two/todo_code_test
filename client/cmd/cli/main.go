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

		fmt.Println("削除・作成・更新・検索・取得・終了のどれかを入力してください。")
		fmt.Scan(&input_value)
		if input_value == "終了" {
			return
		}

		switch input_value {
		case "作成":
			CreateTodo(TodoServiceClient)
		case "取得":
			GetAll(TodoServiceClient)
		case "検索":
			FindSerch(TodoServiceClient)

		}

	}
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
	fmt.Println("期限の月を半角数値で指定してください。")
	fmt.Scanf("月:", &limit_month)
	limit_mouth_as_int = ToInt(limit_month)
	fmt.Println("期限の日を半角数値で入力してください。")
	fmt.Println("日:", limit_day)

	limit_day_as_int = ToInt(limit_day)
	year := time.Now().Local().Year()
	limit_date := time.Date(year, time.Month(limit_mouth_as_int), limit_day_as_int, 0, 0, 0, 0, time.Local)

	res, _ := todo_service_client.CreateTodo(context.Background(),
		connect.NewRequest(&v1.CreateTodoRequest{RequestTodo: &v1.Todo{
			Title:       title,
			Description: descrpition,
			LimitTime:   timestamppb.New(limit_date),
			Status:      v1.Status_INCOMPLETE,
		}}),
	)

	fmt.Println(res)

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
			fmt.Println(err.Error())
			return
		}

		for idx, todo := range res.GetResult() {
			PrintTodo(idx, todo)
		}
		fmt.Println("検索終了しますか？ yes/no")
		fmt.Scan(&input_var)
		if input_var == "yes" {
			return
		}

	}
}
