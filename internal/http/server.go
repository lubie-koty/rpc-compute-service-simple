package http

import (
	"errors"
	"log/slog"
	"net/http"
)

type HTTPServer struct {
	Address string
	Logger  *slog.Logger
	*HTTPService
}

func NewHTTPServer(address string, logger *slog.Logger, httpService *HTTPService) *HTTPServer {
	return &HTTPServer{
		Address:     address,
		Logger:      logger,
		HTTPService: httpService,
	}
}

func NewHTTPHandler(httpService *HTTPService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", httpService.Add)
	mux.HandleFunc("/sub", httpService.Sub)
	mux.HandleFunc("/mul", httpService.Mul)
	mux.HandleFunc("/div", httpService.Div)
	return mux
}

func (s *HTTPServer) Serve() error {
	server := &http.Server{
		Addr:    s.Address,
		Handler: NewHTTPHandler(s.HTTPService),
	}

	s.Logger.Info("HTTP server starting", "server address", server.Addr)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	s.Logger.Info("HTTP server stopped", "server address", server.Addr)
	return nil
}
