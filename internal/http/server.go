package http

import (
	"fmt"
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core"
)

type HttpServer struct {
	core.Server
	HttpService
}

func NewHttpServer(server core.Server, httpService HttpService) *HttpServer {
	return &HttpServer{
		Server:      server,
		HttpService: httpService,
	}
}

func NewHttpHandler(httpService HttpService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", httpService.Add)
	mux.HandleFunc("/sub", httpService.Sub)
	mux.HandleFunc("/mul", httpService.Mul)
	mux.HandleFunc("/div", httpService.Div)
	return mux
}

func (s *HttpServer) Serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Server.Host, s.Server.Port),
		Handler: NewHttpHandler(s.HttpService),
	}

	s.Server.Logger.Info("HTTP server starting", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	s.Server.Logger.Info("HTTP server stopped", server.Addr)
	return nil
}
