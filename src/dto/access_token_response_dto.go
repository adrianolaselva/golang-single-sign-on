package dto

type AccessTokenResponseDTO struct {
	TokenType string `json:"token_type,omitempty"`
	ExpiresIn int64 `json:"expires_in,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
}