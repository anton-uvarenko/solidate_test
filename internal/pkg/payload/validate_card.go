package payload

type ValidateCardResponse struct {
	Valid string        `json:"valid"`
	Error ErrorResponse `json:"error,omitempty"`
}

type ValidCardRequest struct {
	CardNumber      string `json:"cardNumber"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
}
