package grpc_server

import (
	"context"

	pb "api/internal/grpc_gen/todo/v1"
)

// この型にGrpc Serverは依存する
type IWorkflow[InsertType, ResultType any] interface {
	Start(ctx context.Context, arg InsertType)
	Result() (ResultType, error)
}

type TodoCreateServiceServer struct {
	// 注入されたワークフロー　：　インターフェースで抽象度を高める
	injected_workflow IWorkflow[*pb.CreateTodoRequest, *pb.CreateTodoResponse]
}

func NewTodoCreateServiceServer(di_workflow IWorkflow[*pb.CreateTodoRequest, *pb.CreateTodoResponse]) *TodoCreateServiceServer {

	return &TodoCreateServiceServer{
		injected_workflow: di_workflow,
	}
}

func (tcss *TodoCreateServiceServer) Connect(ctx context.Context, request *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	workflow := tcss.injected_workflow
	workflow.Start(ctx, request)
	return workflow.Result()
}
