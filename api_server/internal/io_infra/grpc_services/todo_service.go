package grpc_services

import (
	repo "api/internal/domain/repository/todo_repository"
	todo_use_case "api/internal/domain/use_cases/todo_usecase"
	values "api/internal/domain/values/todo_values"
	v1 "api/internal/grpc_gen/todo/v1"
	"context"
	"log/slog"
	"sync"

	"connectrpc.com/connect"
)

type TodoServer struct {
	todo_use_case *todo_use_case.TodoUseCase[*repo.TodoRepository]
}

func NewTodoGrpcServer() *TodoServer {
	todo_repo, err := repo.NewTodoRepostory()
	if err != nil {
		panic(err)
	}
	return &TodoServer{
		todo_use_case: todo_use_case.NewTodoUseCase(todo_repo),
	}
}

func (ts TodoServer) CreateTodo(ctx context.Context, req *connect.Request[v1.CreateTodoRequest]) (*connect.Response[v1.CreateTodoResponse], error) {
	grpc_todo := req.Msg.GetRequestTodo()
	entity_todo, err := GrpcMessageToEntity(grpc_todo)
	if err != nil {
		return connect.NewResponse(&v1.CreateTodoResponse{
			Result: false,
			Error:  err.Error(),
		}), err
	}

	entity_todo_with_id, err := ts.todo_use_case.CreateTodo(ctx, *entity_todo)
	if err != nil {
		return connect.NewResponse(&v1.CreateTodoResponse{
			Result: false,
			Error:  err.Error(),
		}), err
	}

	grpc_todo = EntityToGrpcMessage(entity_todo_with_id)

	return connect.NewResponse(&v1.CreateTodoResponse{
		Result:      true,
		CreatedTodo: grpc_todo,
		Error:       "",
	}), nil
}
func (ts TodoServer) GetAllTodo(ctx context.Context, req *connect.Request[v1.GetALLRequest]) (*connect.Response[v1.TodoListResponse], error) {

	entity_todo_list, err := ts.todo_use_case.GetAllTodo(ctx, req.Msg.GetIsSort())
	if err != nil {
		return connect.NewResponse(&v1.TodoListResponse{
			Result: make([]*v1.Todo, 0),
			Error:  err.Error(),
		}), err
	}
	var grpc_todos []*v1.Todo
	for _, todo := range entity_todo_list {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}
	return connect.NewResponse(&v1.TodoListResponse{
		Result: grpc_todos,
		Error:  "",
	}), nil
}

func (ts TodoServer) FindTodo(ctx context.Context, stream *connect.BidiStream[v1.SearchRequest, v1.TodoListResponse]) error {

	Request_channel := make(chan *v1.SearchRequest, 100)
	Response_channel := make(chan *v1.TodoListResponse, 100)
	var request_err error
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() { //受信専用のイベントループを走らせるゴルーチン
		defer wg.Done()
		for {
			req, err := stream.Receive()
			if err != nil {
				slog.Log(ctx, slog.LevelError, err.Error())
				request_err = err
				close(Request_channel)
				return
			}
			Request_channel <- req
		}
	}()

	go func() { // ユースケースにリクエストを渡してレスポンスを送信専用チャネルに渡すためのゴルーチン
		defer wg.Done()
		for req := range Request_channel {
			todos, err := ts.todo_use_case.FindAll(ctx, req.GetQuery(), req.GetIsSort())
			var grpc_todos []*v1.Todo
			for _, todo := range todos {
				grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
			}
			var err_msg string
			if err != nil {
				err_msg = err.Error()
			}
			res := &v1.TodoListResponse{
				Result: grpc_todos,
				Error:  err_msg,
			}
			err_msg = ""
			Response_channel <- res
		}
		close(Response_channel) //送信専用ゴルーチンのイベントループを終了させるためのclose
	}()

	go func() { //送信専用のゴルーチン
		defer wg.Done()
		for res := range Response_channel {
			stream.Send(res)
		}
	}()

	wg.Wait()

	return request_err
}

func (ts TodoServer) DeleteTodo(ctx context.Context, req *connect.Request[v1.DeleteTodoRequest]) (*connect.Response[v1.DeleteTodoResponse], error) {
	target_id_as_1nt32 := req.Msg.GetId()
	target_id, err := values.NewTaskId(int(target_id_as_1nt32))

	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error())
		return connect.NewResponse(&v1.DeleteTodoResponse{Result: false, AtherTodo: []*v1.Todo{}}), err
	}
	entiry_todos, ok, err := ts.todo_use_case.DeleteTodo(ctx, target_id)
	grpc_todos := make([]*v1.Todo, 0)

	for _, todo := range entiry_todos {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}

	return connect.NewResponse(&v1.DeleteTodoResponse{
		Result:    ok,
		AtherTodo: grpc_todos,
		Error:     err.Error(),
	}), err
}

func (ts TodoServer) UpdateTodo(ctx context.Context, req *connect.Request[v1.UpdateTodoRequest]) (*connect.Response[v1.UpdateTodoResponse], error) {
	grpc_todo := req.Msg.GetTodo()
	entity_todo, err := GrpcMessageToEntity(grpc_todo)
	if err != nil {
		return connect.NewResponse(&v1.UpdateTodoResponse{Result: false, Err: err.Error()}), err
	}

	id_in_todo, err := ts.todo_use_case.UpdateTodo(ctx, *entity_todo)

	if err != nil {
		return connect.NewResponse(&v1.UpdateTodoResponse{Result: false, Err: err.Error()}), err
	}

	new_grpc_todo := EntityToGrpcMessage(id_in_todo)
	return connect.NewResponse(&v1.UpdateTodoResponse{
		Result: true,
		Todo:   new_grpc_todo,
	}), nil

}
