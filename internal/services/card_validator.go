package services

import (
	"github.com/anton-uvarenko/solidgate_test/internal/pkg/payload"
	"github.com/anton-uvarenko/solidgate_test/internal/pkg/validate/card"
)

type CardValidatorService struct {
}

func NewCardValidatorService() *CardValidatorService {
	return &CardValidatorService{}
}

func (s *CardValidatorService) IsCardValid(payload payload.ValidCardRequest) (bool, error) {
	return card.IsCardValid(payload)
}
