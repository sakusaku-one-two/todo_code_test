package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository/todo_repository"
	grpc_connection "api/internal/grpc_gen/todo/v1"
	models "api/internal/io_infra/database/models"
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
	new_todo := GrpcMessageToEntity(request_todo)

	if created_todo, ok := tuc.repository.Create(new_todo); ok {
		return &grpc_connection.CreateTodoResponse{Result: ok, CreatedTodo: EntityToGrpcMessage(created_todo)}
	}

	return &grpc_connection.CreateTodoResponse{Result: false, CreatedTodo: request_todo}
}

func (tuc *TodoUseCase[repoType]) GetAllTodo(req *grpc_connection.GetALLRequest) *grpc_connection.TodoListResponse {

	todos, err := tuc.repository.GetAll()
	if err != nil {
		return &grpc_connection.CreateTodoResponse{}
	}

}

func EntityToGrpcMessage(grpc_todo entity.Todo) *grpc_connection.Todo {

}

func GrpcMessageToEntity(grpc_todo *grpc_connection.Todo) entity.Todo {

}

// service TodoService {
//     rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
//     rpc GetAllTodo(GetALLRequest) returns(TodoList){};
//     rpc FindTodo(stream SearchRequest) returns(stream TodoList){};
// }
