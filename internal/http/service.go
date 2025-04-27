package http

import (
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core"
)

type HttpService struct {
	service core.MathService
}

func NewHttpService(service core.MathService) *HttpService {
	return &HttpService{
		service: service,
	}
}

func (h *HttpService) Add(w http.ResponseWriter, r *http.Request) {
	return h.service.Add(1, 1)
}

func (h *HttpService) Sub(w http.ResponseWriter, r *http.Request) {
	return h.service.Sub(1, 1)
}

func (h *HttpService) Mul(w http.ResponseWriter, r *http.Request) {
	return h.service.Mul(1, 1)
}

func (h *HttpService) Div(w http.ResponseWriter, r *http.Request) {
	return h.service.Div(1, 1)
}
