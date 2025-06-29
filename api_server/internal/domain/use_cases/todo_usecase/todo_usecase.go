package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository"
	values "api/internal/domain/values/todo_values"
	v1 "api/internal/grpc_gen/todo/v1"
	models "api/internal/io_infra/database/models"
	"context"
	"log/slog"
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

func (tuc *TodoUseCase[repoType]) CreateTodo(ctx context.Context, req *v1.CreateTodoRequest) (*v1.CreateTodoResponse, error) {

	request_todo := req.GetRequestTodo()
	title, err := values.NewTitle(request_todo.Title)
	if err != nil {
		return &v1.CreateTodoResponse{
			Error: err.Error(),
		}, err
	}
	limit, err := values.NewLimit(request_todo.LimitTime.AsTime())
	if err != nil {
		return &v1.CreateTodoResponse{
			Error: err.Error(),
		}, err
	}
	description, _ := values.NewDescription(request_todo.Description)
	status, _ := values.GetTodoStatus(int(request_todo.Status))

	entity_todo := entity.Todo{
		Title:       title,
		Description: description,
		Limit:       limit,
		Status:      status,
	}

	created_todo, err := tuc.repository.Create(ctx, entity_todo)
	if err != nil {
		return &v1.CreateTodoResponse{
			Result:      false,
			CreatedTodo: EntityToGrpcMessage(created_todo),
			Error:       err.Error()}, err
	}

	return &v1.CreateTodoResponse{Result: true, CreatedTodo: EntityToGrpcMessage(created_todo), Error: ""}, nil
}

func (tuc *TodoUseCase[repoType]) GetAllTodo(ctx context.Context, req *v1.GetALLRequest) (*v1.TodoListResponse, error) {

	todos, err := tuc.repository.GetAll(ctx)
	if err != nil {
		return &v1.TodoListResponse{Error: err.Error()}, err
	}

	var grpc_todos []*v1.Todo
	for _, todo := range todos {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}
	return &v1.TodoListResponse{Result: grpc_todos, Error: ""}, nil
}

func (tuc *TodoUseCase[repoType]) DeleteTodo(ctx context.Context, req *v1.DeleteTodoRequest) (*v1.DeleteTodoResponse, error) {

	id := int(req.Id)
	task_id, err := values.NewTaskId(id)
	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error()) //タスクIDのdomaiロジックから弾かれたケース
		return &v1.DeleteTodoResponse{
			Result: false,
			Error:  err.Error(),
		}, err
	}

	if ok, err := tuc.repository.Delete(ctx, task_id); !ok {
		return &v1.DeleteTodoResponse{
			Result: false,
			Error:  err.Error(),
		}, err
	}

	todos, err := tuc.repository.GetAll(ctx)
	if err != nil {
		return &v1.DeleteTodoResponse{Result: false}, err
	}
	var result []*v1.Todo
	for _, todo := range todos {
		result = append(result, EntityToGrpcMessage(todo))
	}

	return &v1.DeleteTodoResponse{
		Result:    true,
		AtherTodo: result,
	}, nil
}

func (tuc *TodoUseCase[repoType]) FindAll(ctx context.Context, req *v1.SearchRequest) (*v1.TodoListResponse, error) {
	query_string := req.GetQuery()

	todos, err := tuc.repository.FindAll(ctx, query_string)
	if err != nil {
		return &v1.TodoListResponse{
			Result: make([]*v1.Todo, 0),
			Error:  err.Error(),
		}, nil
	}

	var grpc_todos []*v1.Todo
	for _, todo := range todos {
		grpc_todos = append(grpc_todos, EntityToGrpcMessage(todo))
	}

	return &v1.TodoListResponse{
		Result: grpc_todos,
		Error:  "",
	}, nil
}

func (tuc *TodoUseCase[repoType]) UpdateTodo(ctx context.Context, req *v1.UpdateTodoRequest) (*v1.UpdateTodoResponse, error) {
	todo := req.GetTodo()
	entity_todo, err := GrpcMessageToEntity(todo)

	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error()+" at Grpc Todo => entity todo trancelate funcition")
		return &v1.UpdateTodoResponse{Result: false, Err: err.Error()}, err
	}

	_, err = tuc.repository.Update(ctx, *entity_todo)
	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error()+" at repository.Update")
		return &v1.UpdateTodoResponse{Result: false, Err: err.Error()}, err
	}

	return &v1.UpdateTodoResponse{Result: true, Err: ""}, nil
}
