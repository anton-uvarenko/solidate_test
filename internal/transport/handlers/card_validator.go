package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anton-uvarenko/solidgate_test/internal/pkg/payload"
	"net/http"
)

type CardValidatorHandler struct {
	service iCardValidatorService
}

type iCardValidatorService interface {
	IsCardValid(payload payload.ValidCardRequest) (bool, error)
}

func NewCardValidatorHandler(service iCardValidatorService) *CardValidatorHandler {
	return &CardValidatorHandler{
		service: service,
	}
}

func (h *CardValidatorHandler) IsCardValid(w http.ResponseWriter, r *http.Request) {
	pl := &payload.ValidCardRequest{}
	err := json.NewDecoder(r.Body).Decode(pl)
	if err != nil {
		http.Error(w, fmt.Errorf("%w: %w", payload.ErrBadRequest, err).Error(), http.StatusBadRequest)
		return
	}

	isValid, err := h.service.IsCardValid(*pl)
	response := &payload.ValidateCardResponse{
		Valid: isValid,
	}

	if err != nil {
		response.Error = err.(*payload.Error)
	}

	json.NewEncoder(w).Encode(response)
}
