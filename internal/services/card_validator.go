package services

import (
	"solidgate_test/internal/pkg/payload"
	"solidgate_test/internal/pkg/validate"
)

type CardValidatorService struct {
}

func NewCardValidatorService() *CardValidatorService {
	return &CardValidatorService{}
}

func (s *CardValidatorService) IsCardValid(payload payload.ValidCardRequest) (bool, error) {
	return validate.IsCardValid(payload)
}
