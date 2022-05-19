package storefrontlayout

import "github.com/mkamadeus/iot-smart-retail/service/storefront"

type Handler struct {
	Service *storefront.Service
}

func New(service *storefront.Service) *Handler {
	return &Handler{
		Service: service,
	}
}
