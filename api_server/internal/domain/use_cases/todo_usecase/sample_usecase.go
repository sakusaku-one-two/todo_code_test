package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository/todo_repository"
	values "api/internal/domain/values/todo_values"
	grpc_connection "api/internal/grpc_gen/todo/v1"
	models "api/internal/io_infra/database/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo]] struct {
	repository repoType
}

func NewTodoUseCase[repoType repo.IRepository[models.Todo, entity.Todo]](repo repoType) *TodoUseCase[repoType] {
	todo_usecase := &TodoUseCase[repoType]{
		repository: repo,
	}
	return todo_usecase
}

func (tuc *TodoUseCase[repoType]) CreateTodo(req *grpc_connection.CreateTodoRequest) *grpc_connection.CreateTodoResponse {

	request_todo := req.GetRequestTodo()
	entity_todo := GrpcMessageToEntity(request_todo)

	if created_todo, ok := tuc.repository.Create(entity_todo); ok {
		return &grpc_connection.CreateTodoResponse{Result: ok, CreatedTodo: EntityToGrpcMessage(created_todo)}
	}

	return &grpc_connection.CreateTodoResponse{Result: false, CreatedTodo: request_todo}
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

	if ok := tuc.repository.Delete(task_id); ok == false {
		return &grpc_connection.DeleteTodoResponse{
			Result: false,
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
		Id:          entity_todo.Id.GetValue(),
		Title:       entity_todo.Title.GetValue(),
		Description: entity_todo.Description.GetValue(),
		LimitTime:   timestamppb.New(entity_todo.Limit.GetValue()),
		CreatedAt:   timestamppb.New(entity_todo.Created_at),
		UpdatedAt:   timestamppb.New(entity_todo.Update_at),
		Status:      entity_todo.Status,
	}
}

func GrpcMessageToEntity(grpc_todo *grpc_connection.Todo) entity.Todo {

}

// service TodoService {
//     rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
//     rpc GetAllTodo(GetALLRequest) returns(TodoList){};
//     rpc FindTodo(stream SearchRequest) returns(stream TodoList){};
// }
