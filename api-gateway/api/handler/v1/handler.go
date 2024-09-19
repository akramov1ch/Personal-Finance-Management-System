package v1

import (
	"api-gateway/service"
)

type handlerV1 struct {
	service service.IClients
}

type HandlerConfig struct {
	Service service.IClients
}

func NewHandlerV1(h *HandlerConfig) *handlerV1 {
	return &handlerV1{
		service: h.Service,
	}
}
