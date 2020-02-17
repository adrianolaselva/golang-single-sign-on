package contracts

type Token struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken string `json:"id_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int32 `json:"expires_in"`
}
