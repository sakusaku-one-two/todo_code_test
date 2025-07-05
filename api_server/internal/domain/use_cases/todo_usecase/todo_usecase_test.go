package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	repo "api/internal/domain/repository/todo_repository"
	values "api/internal/domain/values/todo_values"
	"context"
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

// func Test_todo_create(t *testing.T) {
// 	ctx := t.Context()
// 	create_todo_res, _ := todo_use_case.CreateTodo(ctx, &grpc_connection.CreateTodoRequest{
// 		RequestTodo: generate_grpc_todo(1, "new_todo_sample", "description_new_todo", time.Now().Add(100*time.Minute)),
// 	})

// 	if create_todo_res.Error != "" {
// 		t.Fatalf("create todo err => %s", create_todo_res.GetError())
// 	}

// 	fmt.Println(
// 		"new todo create complete!!",
// 		create_todo_res,
// 	)

// }

// func Test_todo_getall(t *testing.T) {
// 	ctx := t.Context()
// 	res, _ := todo_use_case.GetAllTodo(ctx, &grpc_connection.GetALLRequest{
// 		Request: "",
// 		IsSort:  true,
// 	})
// 	for _, todo := range res.GetResult() {
// 		fmt.Println(todo)
// 	}

// }

// func Test_delete_from_use_case(t *testing.T) {
// 	new_todo := generate_grpc_todo(12,"sample","sample_description",time.Now().Add(15*time.Hour))
// 	entiry_todo,err := ConvertModel[v1.Todo,entity.Todo](new_todo)
// }

// func Test_todo_findall(t *testing.T) {
// 	res := todo_use_case.
// }

func time_(t time.Duration) values.Limit {
	return_value, _ := values.NewLimit(time.Now().Add(t))
	return return_value
}

func CreateTodo() []entity.Todo {
	return []entity.Todo{
		entity.Todo{Limit: time_(100 * time.Hour)},
		entity.Todo{Limit: time_(10 * time.Hour)},
		entity.Todo{Limit: time_(170 * time.Hour)},
		entity.Todo{Limit: time_(150 * time.Hour)},
	}

}

func Test_avarage(t *testing.T) {
	todos := CreateTodo()
	for i, todo := range todos {
		fmt.Println(i, todo.Limit.GetValue())
	}
	ctx := context.Background()
	sort := NewSort(RecSort)
	reuslt := sort.Execute(ctx, todos)
	for idx, todo := range reuslt {
		fmt.Println(idx, todo.Limit.GetValue())
	}

}
