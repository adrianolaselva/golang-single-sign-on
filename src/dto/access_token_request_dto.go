package dto

type AccessTokenRequestDTO struct {
	GrantType string `json:"grant_type,omitempty" schema:"grant_type,omitempty"`
	ClientID string `json:"client_id,omitempty" schema:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty" schema:"client_secret,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty" schema:"redirect_uri,omitempty"`
	RefreshType string `json:"refresh_token,omitempty" schema:"refresh_token,omitempty"`
	Scope string `json:"scope,omitempty" schema:"scope,omitempty"`
	Username string `json:"username,omitempty" schema:"username,omitempty"`
	Password string `json:"password,omitempty" schema:"password,omitempty"`
}