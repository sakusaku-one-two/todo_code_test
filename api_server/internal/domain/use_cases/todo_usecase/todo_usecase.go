package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository"
	values "api/internal/domain/values/todo_values"
	grpc_connection "api/internal/grpc_gen/todo/v1"
	models "api/internal/io_infra/database/models"
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
	title, err := values.NewTitle(request_todo.Title)
	if err != nil {
		return &grpc_connection.CreateTodoResponse{
			Error: err.Error(),
		}
	}
	limit, err := values.NewLimit(request_todo.LimitTime.AsTime())
	if err != nil {
		return &grpc_connection.CreateTodoResponse{
			Error: err.Error(),
		}
	}
	description, _ := values.NewDescription(request_todo.Description)
	status, _ := values.GetTodoStatus(int(request_todo.Status))

	entity_todo := entity.Todo{
		Title:       title,
		Description: description,
		Limit:       limit,
		Status:      status,
	}

	created_todo, err := tuc.repository.Create(entity_todo)
	if err != nil {
		return &grpc_connection.CreateTodoResponse{
			Result:      false,
			CreatedTodo: EntityToGrpcMessage(created_todo),
			Error:       err.Error()}
	}

	return &grpc_connection.CreateTodoResponse{Result: true, CreatedTodo: EntityToGrpcMessage(created_todo), Error: ""}
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
	task_id, err := values.NewTaskId(id)

	if ok, err := tuc.repository.Delete(task_id); !ok {
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

func (tuc *TodoUseCase[repoType]) FindAll(req *grpc_connection.SearchRequest) *grpc_connection.TodoListResponse {
	query_string := req.GetQuery()

	todos, err := tuc.repository.FindAll(query_string)
	if err != nil {
		return &grpc_connection.TodoListResponse{
			Result: make([]*grpc_connection.Todo, 0),
			Error:  err.Error(),
		}
	}

	var grpc_todos []*grpc_connection.Todo
	for _, todo := range todos {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}

	return &grpc_connection.TodoListResponse{
		Result: grpc_todos,
		Error:  "",
	}
}

// service TodoService {
//     rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
//     rpc GetAllTodo(GetALLRequest) returns(TodoList){};
//     rpc FindTodo(stream SearchRequest) returns(stream TodoList){};
// }
