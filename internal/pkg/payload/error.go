package payload

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
