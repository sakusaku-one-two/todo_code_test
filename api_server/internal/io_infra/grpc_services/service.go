package grpc_services

import (
	pb "api/internal/grpc_gen/todo/v1"
	"context"

	"connectrpc.com/connect"
)

type DependencyInjecteWorkfrow = func(ctx *context.Context, req *connect.Request[pb.CreateTodoRequest]) (*connect.Response[pb.CreateTodoResponse], error)

var (
	TODO_SERVER = &TodoServer[func()]{}
)

type TodoServer[inject_funciton DependencyInjecteWorkfrow] struct {
	injected_workfrow inject_funciton
}

func (ts *TodoServer[inject_funciton]) CreateTodo(
	ctx context.Context,
	req *connect.Request[pb.CreateTodoRequest],
) (*connect.Response[pb.CreateTodoResponse], error) {
	return ts.injected_workfrow(ctx, req)
}
