package main

import (
	hadler "api/internal/handler"
	server_config "api/internal/io_infra/config/server_config"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	//　実行のエントリーポイント
	mux := http.NewServeMux()
	hadler.SetUpHandler(mux)
	http.ListenAndServe(
		server_config.GetServerAddressAndPort(),
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
