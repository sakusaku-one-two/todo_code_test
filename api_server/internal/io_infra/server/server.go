package server

import (
	"net/http"
)

func NewServer() (*http.ServeMux, error) {
	new_mux := http.ServeMux{}
	return &new_mux, nil
}
