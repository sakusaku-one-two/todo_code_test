package main

import (
	hadler "api/internal/handler"
	server_config "api/internal/io_infra/config"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	//　実行のエントリーポイント
	mux := http.NewServeMux()
	hadler.SetUpHandler(mux)
	http.ListenAndServe(
		fmt.Printf("%s:%s", server_config.SERVER_HOST, server_config.SERVER_PORT),
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
