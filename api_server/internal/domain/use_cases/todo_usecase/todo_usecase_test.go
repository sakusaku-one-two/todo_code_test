package todo_usecase

import (
	repo "api/internal/domain/repository/todo_repository"

	"fmt"
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

var todo_use_case *TodoUseCase[*repo.TodoRepository] = nil

func init() {
	todo_repo, err := repo.NewTodoRepostory()
	if err != nil {
		return
	}
	todo_use_case = NewTodoUseCase(todo_repo)
}

func Test_todo_create(t *testing.T) {
	ctx := t.Context()
	create_todo_res := todo_use_case.CreateTodo(ctx, &grpc_connection.CreateTodoRequest{
		RequestTodo: generate_grpc_todo(1, "new_todo_sample", "description_new_todo", time.Now().Add(100*time.Minute)),
	})

	if create_todo_res.Error != "" {
		t.Fatalf("create todo err => %s", create_todo_res.GetError())
	}

	fmt.Println(
		"new todo create complete!!",
		create_todo_res,
	)

}

func Test_todo_getall(t *testing.T) {
	ctx := t.Context()
	res := todo_use_case.GetAllTodo(ctx, &grpc_connection.GetALLRequest{
		Request: "",
		IsSort:  true,
	})
	for _, todo := range res.GetResult() {
		fmt.Println(todo)
	}

}

// func Test_todo_findall(t *testing.T) {
// 	res := todo_use_case.
// }
