package payload

type ValidateCardResponse struct {
	Valid bool   `json:"valid"`
	Error *Error `json:"error,omitempty"`
}

type ValidCardRequest struct {
	CardNumber      string `json:"cardNumber"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
}
