package handler

import (
	connect "api/internal/grpc_gen/todo/v1/todo_v1connect"
	grpc_service "api/internal/io_infra/grpc_services"
	"net/http"
)

func SetUpHandler(mux *http.ServeMux) {
	path, handler := connect.NewTodoServiceHandler(&grpc_service.TodoServer{})
	mux.Handle(path, handler)
}
