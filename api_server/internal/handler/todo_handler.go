package handler

import (
	connect "api/internal/grpc_gen/todo/v1/todov1connect"
	grpc_service "api/internal/io_infra/grpc_services"
	"net/http"
)

func SetUpHandler(mux *http.ServeMux) {
	path, handler := connect.NewTodoServiceHandler(grpc_service.NewTodoGrpcServer())
	mux.Handle(path, handler)
}
