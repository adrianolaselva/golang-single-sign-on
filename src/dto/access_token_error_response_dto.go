package dto

type AccessTokenErrorResponseDTO struct {
	Status int8 `json:"status,omitempty"`
	Error string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}