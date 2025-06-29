package todo_repository

import (
	entity "api/internal/domain/entitys/todo_entity"
	values "api/internal/domain/values/todo_values"
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_repo_update(t *testing.T) {
	repo, _ := NewTodoRepostory()

	ctx := context.Background()
	title, _ := values.NewTitle("sakusaku")
	desc, _ := values.NewDescription("sakuksau desc")
	limit, _ := values.NewLimit(time.Now())
	status, _ := values.GetTodoStatus(0)
	todo := entity.Todo{
		Title:       title,
		Description: desc,
		Limit:       limit,
		Status:      status,
	}

	updated_todo, err := repo.Update(ctx, todo)
	fmt.Println(updated_todo, err, ctx)
}

func Test_repo(t *testing.T) {
	repo, err := NewTodoRepostory()
	if err != nil || repo == nil {
		return
	}
	ctx := context.Background()
	// title, _ := values.NewTitle("sakusaku")
	// desc, _ := values.NewDescription("sakuksau desc")
	// limit, _ := values.NewLimit(time.Now())
	// fmt.Println(limit)
	// status, _ := values.GetTodoStatus(0)
	// todo := entity.Todo{
	// 	Title:       title,
	// 	Description: desc,
	// 	Limit:       limit,
	// 	Status:      status,
	// }
	// repo.Create(ctx, todo)

	data, err := repo.FindAll(
		ctx,
		"n",
	)

	fmt.Println("find data => ", data, err)

}
