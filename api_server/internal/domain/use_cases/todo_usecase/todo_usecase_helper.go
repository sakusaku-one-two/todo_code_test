package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	values "api/internal/domain/values/todo_values"
	grpc_connection "api/internal/grpc_gen/todo/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EntityToGrpcMessage(entity_todo entity.Todo) *grpc_connection.Todo {

	return &grpc_connection.Todo{
		Id:          int_to_int32(entity_todo.Id.GetValue()),
		Title:       entity_todo.Title.GetValue(),
		Description: entity_todo.Description.GetValue(),
		LimitTime:   timestamppb.New(entity_todo.Limit.GetValue()),
		CreatedAt:   timestamppb.New(entity_todo.Created_at),
		UpdatedAt:   timestamppb.New(entity_todo.Update_at),
		Status:      grpc_connection.Status(values.GetTodoStatusNumber(entity_todo.Status)),
		IsActivate:  entity_todo.Is_activate,
	}
}

func GrpcMessageToEntity(grpc_todo *grpc_connection.Todo) (*entity.Todo, error) {

	title, err := values.NewTitle(grpc_todo.Title)
	if err != nil {
		return nil, err
	}

	Description, err := values.NewDescription(grpc_todo.Description)
	if err != nil {
		return nil, err
	}

	status, err := values.GetTodoStatus(int(grpc_todo.Status.Number()))
	if err != nil {
		return nil, err
	}

	id, err := values.NewTaskId(int32_to_int(grpc_todo.Id))
	if err != nil {
		return nil, err
	}

	limit, err := values.NewLimit(grpc_todo.LimitTime.AsTime())
	if err != nil {
		return nil, err
	}

	return &entity.Todo{
		Id:          id,
		Title:       title,
		Description: Description,
		Limit:       limit,
		Status:      status,
		Is_activate: grpc_todo.IsActivate,
		Update_at:   grpc_todo.UpdatedAt.AsTime(),
		Created_at:  grpc_todo.CreatedAt.AsTime(),
	}, nil

}

func int32_to_int(target *int32) int {
	return int(*target)
}

func int_to_int32(target int) *int32 {
	result := int32(target)
	return &result
}
