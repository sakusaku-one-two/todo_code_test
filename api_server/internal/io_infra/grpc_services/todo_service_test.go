package grpc_services

import (
	todov1 "api/internal/grpc_gen/todo/v1"
	"context"
	"fmt"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMain(t *testing.T) {
	if t.Run("create todo!", Test_create) && t.Run("get all", Test_getall) && t.Run("all delete!!", Test_delete) {
		t.Skip()
	} else {
		t.Error("test todo grpc server failed")
	}
}

func Test_getall(t *testing.T) {
	todo_server := NewTodoGrpcServer()
	ctx := context.Background()
	res, err := todo_server.GetAllTodo(ctx, connect.NewRequest(&todov1.GetALLRequest{
		Request: "",
		IsSort:  false,
	}))

	if err != nil {
		return
	}

	fmt.Println("todo server(grpc) get test DONE!!", res)

}

func Test_create(t *testing.T) {
	todo_server := NewTodoGrpcServer()
	ctx := t.Context()
	res, err := todo_server.CreateTodo(ctx, connect.NewRequest(&todov1.CreateTodoRequest{
		RequestTodo: &todov1.Todo{
			Title:       "grpce_server_create_test",
			Description: "grpc_server_create_test_description",
			LimitTime:   timestamppb.New(time.Now().Add(100 * time.Minute)),
			Status:      todov1.Status_INCOMPLETE,
		},
	}))

	fmt.Println("create =>", res, err)
}

func Test_delete(t *testing.T) {
	todo_server := NewTodoGrpcServer()
	ctx := t.Context()
	res, err := todo_server.GetAllTodo(ctx, connect.NewRequest(&todov1.GetALLRequest{
		Request: "",
		IsSort:  false,
	}))

	if err != nil {
		return
	}
	fmt.Println("to delete")

	for _, todo := range res.Msg.Result {
		fmt.Println(todo)
		res, _ := todo_server.DeleteTodo(ctx, connect.NewRequest(&todov1.DeleteTodoRequest{
			Id: *todo.Id,
		}))
		fmt.Println("DELETED TODO => ", res)
	}
}
