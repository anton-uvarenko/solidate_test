package services

type Service struct {
	CardValidatorService *CardValidatorService
}

func NewService() *Service {
	return &Service{
		CardValidatorService: NewCardValidatorService(),
	}
}
