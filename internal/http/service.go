package http

import (
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core/types"
	"github.com/lubie-koty/rpc-compute-service-simple/internal/util"
)

type HttpService struct {
	service types.MathService
}

func NewHttpService(service types.MathService) *HttpService {
	return &HttpService{
		service: service,
	}
}

type OperationRequest struct {
	FirstNumber  int32 `json:"first_number" validate:"required"`
	SecondNumber int32 `json:"second_number" validate:"required"`
}

type OperationResponse struct {
	Result int32 `json:"result"`
}

func (h *HttpService) Add(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Add(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HttpService) Sub(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Sub(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HttpService) Mul(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Mul(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (h *HttpService) Div(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := h.service.Div(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}
