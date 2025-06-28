package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository"

	// "api/internal/domain/values"
	values "api/internal/domain/values/todo_values"
	grpc_connection "api/internal/grpc_gen/todo/v1"
	models "api/internal/io_infra/database/models"
	"context"
	"log/slog"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo, values.TaskId[int]]] struct {
	repository repoType
}

func NewTodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo, values.TaskId[int]]](repo repoType) *TodoUseCase[repoType] {
	todo_usecase := &TodoUseCase[repoType]{
		repository: repo,
	}
	return todo_usecase
}

func (tuc *TodoUseCase[repoType]) CreateTodo(req *grpc_connection.CreateTodoRequest) *grpc_connection.CreateTodoResponse {

	request_todo := req.GetRequestTodo()
	entity_todo, err := GrpcMessageToEntity(request_todo)

	if err != nil {
		ctx := context.Background()
		slog.Log(ctx, slog.LevelInfo, err.Error())
		return &grpc_connection.CreateTodoResponse{Error: ""}
	}

	if created_todo, ok := tuc.repository.Create(*entity_todo); ok {
		return &grpc_connection.CreateTodoResponse{Result: ok, CreatedTodo: EntityToGrpcMessage(created_todo), Error: ""}
	}

	return &grpc_connection.CreateTodoResponse{Result: false, CreatedTodo: request_todo, Error: "申し訳ありません。作成に失敗いたしました。"}
}

func (tuc *TodoUseCase[repoType]) GetAllTodo(req *grpc_connection.GetALLRequest) *grpc_connection.TodoListResponse {

	todos, err := tuc.repository.GetAll()
	if err != nil {
		return &grpc_connection.TodoListResponse{Error: err.Error()}
	}

	var grpc_todos []*grpc_connection.Todo
	for _, todo := range todos {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}
	return &grpc_connection.TodoListResponse{Result: grpc_todos, Error: ""}
}

func (tuc *TodoUseCase[repoType]) DeleteTodo(req *grpc_connection.DeleteTodoRequest) *grpc_connection.DeleteTodoResponse {
	id := int(req.Id)
	task_id, _ := values.NewTaskId(id)

	if ok, err := tuc.repository.Delete(task_id); ok == false {
		return &grpc_connection.DeleteTodoResponse{
			Result: false,
			Error:  err.Error(),
		}
	}

	todos, err := tuc.repository.GetAll()
	if err != nil {
		return &grpc_connection.DeleteTodoResponse{Result: false}
	}
	var result []*grpc_connection.Todo
	for _, todo := range todos {
		result = append(result, EntityToGrpcMessage(todo))
	}

	return &grpc_connection.DeleteTodoResponse{
		Result:    true,
		AtherTodo: result,
	}
}

func EntityToGrpcMessage(entity_todo entity.Todo) *grpc_connection.Todo {

	return &grpc_connection.Todo{
		Id:          int_to_int32(entity_todo.Id.GetValue()),
		Title:       entity_todo.Title.GetValue(),
		Description: entity_todo.Description.GetValue(),
		LimitTime:   timestamppb.New(entity_todo.Limit.GetValue()),
		CreatedAt:   timestamppb.New(entity_todo.Created_at),
		UpdatedAt:   timestamppb.New(entity_todo.Update_at),
		Status:      grpc_connection.Status(values.GetTodoStatusNumber(entity_todo.Status)),
		IsActivate:  entity_todo.Is_activate,
	}
}

func int_to_int32(target int) *int32 {
	result := int32(target)
	return &result
}

func GrpcMessageToEntity(grpc_todo *grpc_connection.Todo) (*entity.Todo, error) {
	title, err := values.NewTitle(grpc_todo.Title)
	if err != nil {
		return nil, err
	}

	Description, err := values.NewDescription(grpc_todo.Description)
	if err != nil {
		return nil, err
	}

	status, err := values.GetTodoStatus(int(grpc_todo.Status.Number()))
	if err != nil {
		return nil, err
	}
	id, err := values.NewTaskId(int32_to_int(grpc_todo.Id))
	if err != nil {
		return nil, err
	}

	limit, err := values.NewLimit(grpc_todo.LimitTime.AsTime())
	if err != nil {
		return nil, err
	}

	return &entity.Todo{
		Id:          id,
		Title:       title,
		Description: Description,
		Limit:       limit,
		Status:      status,
		Is_activate: grpc_todo.IsActivate,
		Update_at:   grpc_todo.UpdatedAt.AsTime(),
		Created_at:  grpc_todo.CreatedAt.AsTime(),
	}, nil

}

func int32_to_int(target *int32) int {
	return int(*target)
}

// service TodoService {
//     rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
//     rpc GetAllTodo(GetALLRequest) returns(TodoList){};
//     rpc FindTodo(stream SearchRequest) returns(stream TodoList){};
// }
