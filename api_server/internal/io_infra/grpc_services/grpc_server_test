package grpc_server

import (
	pb "api/internal/grpc_gen/todo/v1"
	"api/util"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"testing"
)

type SampleWorkFlow struct {
}

func (swf *SampleWorkFlow) Start(ctx context.Context, arg *pb.CreateTodoRequest) {
	if util.IsNil(swf) {
		return
	}
	slog.Info("start!!")
}

func (swf *SampleWorkFlow) Result() (*pb.CreateTodoResponse, error) {
	if util.IsNil(swf) {
		return nil, errors.New("swf is nil")
	}
	return &pb.CreateTodoResponse{
		Result: true,
	}, nil
}

func Test_grpc_createServer(t *testing.T) {
	fmt.Println("test grpc create server")

	sample_wf := &SampleWorkFlow{}
	samle_server := NewTodoCreateServiceServer(sample_wf)
	ctx := context.Background()
	todo := pb.Todo{
		Title:       "sample",
		Description: "sample description",
		Limit:       nil,
	}
	sample_request := pb.CreateTodoRequest{
		NewTodo: &todo,
	}
	responese, err := samle_server.Connect(ctx, &sample_request)
	if err != nil {
		t.Log("ok!")
	}

	t.Log(responese)
}
