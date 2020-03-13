package dto

type AccessTokenErrorResponseDTO struct {
	Error string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}