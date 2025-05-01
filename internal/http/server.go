package http

import (
	"errors"
	"log/slog"
	"net/http"
)

type HttpServer struct {
	Address string
	Logger  *slog.Logger
	*HttpService
}

func NewHttpServer(address string, logger *slog.Logger, httpService *HttpService) *HttpServer {
	return &HttpServer{
		Address:     address,
		Logger:      logger,
		HttpService: httpService,
	}
}

func NewHttpHandler(httpService *HttpService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", httpService.Add)
	mux.HandleFunc("/sub", httpService.Sub)
	mux.HandleFunc("/mul", httpService.Mul)
	mux.HandleFunc("/div", httpService.Div)
	return mux
}

func (s *HttpServer) Serve() error {
	server := &http.Server{
		Addr:    s.Address,
		Handler: NewHttpHandler(s.HttpService),
	}

	s.Logger.Info("HTTP server starting", "server address", server.Addr)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	s.Logger.Info("HTTP server stopped", "server address", server.Addr)
	return nil
}
