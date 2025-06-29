package todo_repository

import (
	"context"
	"fmt"
	"testing"
)

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
