package main

import "net/http"

func main() {
	//　実行のエントリーポイント
	mux := http.ServeMux{}

	server := mux.NewServer()
	server.Server()
}
