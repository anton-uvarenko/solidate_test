package handlers

import "solidgate_test/internal/services"

type Handler struct {
	CardValidatorHandler *CardValidatorHandler
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		CardValidatorHandler: NewCardValidatorHandler(service.CardValidatorService),
	}
}
