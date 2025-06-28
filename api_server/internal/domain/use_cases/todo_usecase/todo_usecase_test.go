package todo_usecase

import (
	repo "api/internal/domain/repository/todo_repository"
	"testing"

	// "api/internal/domain/values"

	grpc_connection "api/internal/grpc_gen/todo/v1"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func generate_grpc_todo(id int32, title string, description string, limittime time.Time) *grpc_connection.Todo {
	return &grpc_connection.Todo{
		Id:          &id,
		Title:       title,
		Description: description,
		LimitTime:   timestamppb.New(limittime),
		Status:      grpc_connection.Status(0),
	}
}

func Test_todo_use_case(t *testing.T) {
	todo_repo, err := repo.NewTodoRepostory()
	if err != nil {
		t.Fatalf("todo create err => %s", err.Error())
	}
	todo_use_case := NewTodoUseCase(todo_repo)

	create_todo_res := todo_use_case.CreateTodo(&grpc_connection.CreateTodoRequest{
		RequestTodo: generate_grpc_todo(1, "new_todo_sample", "description_new_todo", time.Now().Add(100*time.Minute)),
	})

	if create_todo_res.Error != "" {
		t.Fatalf("create todo err => %s", create_todo_res.GetError())
	}

}

// service TodoService {
//     rpc CreateTodo(CreateTodoRequest) returns (CreateTodoResponse){};
//     rpc GetAllTodo(GetALLRequest) returns(TodoList){};
//     rpc FindTodo(stream SearchRequest) returns(stream TodoList){};
// }
