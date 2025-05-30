package http

import (
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core/types"
	"github.com/lubie-koty/rpc-compute-service-simple/internal/util"
)

type HTTPService struct {
	service types.MathService
}

func NewHTTPService(service types.MathService) *HTTPService {
	return &HTTPService{
		service: service,
	}
}

type OperationRequest struct {
	Numbers []float64 `json:"numbers" validate:"required"`
}

type OperationResponse struct {
	Result float64 `json:"result"`
}

func (h *HTTPService) Add(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Add(body.Numbers)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HTTPService) Sub(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Sub(body.Numbers)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HTTPService) Mul(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Mul(body.Numbers)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HTTPService) Div(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Div(body.Numbers)
	util.WriteResponse(w, OperationResponse{Result: result})
}
