package todo_repository

import (
	entity "api/internal/domain/entitys/todo_entity"
	values "api/internal/domain/values/todo_values"
	"fmt"
	"testing"
	"time"
)

func Test_repo(t *testing.T) {
	repo, err := NewTodoRepostory()
	if err != nil || repo == nil {
		return
	}

	title, _ := values.NewTitle("sakusaku")
	desc, _ := values.NewDescription("sakuksau desc")
	limit, _ := values.NewLimit(time.Now())
	fmt.Println(limit)
	status, _ := values.GetTodoStatus(0)
	todo := entity.Todo{
		Title:       title,
		Description: desc,
		Limit:       limit,
		Status:      status,
	}
	repo.Create(todo)

	data, err := repo.FindAll(
		"saku",
	)

	fmt.Println("find => ", data, err)

}
