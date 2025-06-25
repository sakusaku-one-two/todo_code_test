package grpc_services

import (
	v1 "api/internal/grpc_gen/todo/v1"
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
)

type TodoServer struct {
}

func (ts TodoServer) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {

	todo := req.Msg.RequestTodo

	res := connect.NewResponse(&v1.CreateTodoResponse{
		Result:      true,
		CreatedTodo: todo,
	})

	fmt.Println(todo)
	return res, nil
}

func (ts TodoServer) GetAllTodo(ctx context.Context, req *connect.Request[v1.GetALLRequest]) (*connect.Response[v1.TodoListResponse], error) {
	todos := []*v1.Todo{}
	for i := 0; i < 10; i++ {

		temp := &v1.Todo{
			Title:       "inosaku",
			Description: "sakusaku",
		}
		todos = append(todos, temp)
	}

	return connect.NewResponse(&v1.TodoListResponse{Result: todos, Error: ""}), nil
}

func (ts TodoServer) FindTodo(ctx context.Context, stream *connect.BidiStream[v1.SearchRequest, v1.TodoListResponse]) error {

	go func() {
		for {
			req, err := stream.Receive()
			if err != nil {
				return
			}
			fmt.Println(req.Query)
		}
	}()

	Ticker := time.NewTicker(10 * time.Second)

	for {
		<-Ticker.C

		todos := []*v1.Todo{}
		for i := 0; i < 10; i++ {

			temp := &v1.Todo{
				Title:       "inosaku",
				Description: "sakusaku",
			}
			todos = append(todos, temp)
		}

		stream.Send(&v1.TodoListResponse{Result: todos, Error: ""})
	}

}

func (TodoServer) DeleteTodo(ctx context.Context, req *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error) {
	return nil, nil
}
