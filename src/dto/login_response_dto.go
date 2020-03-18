package dto


type LoginResponseDTO struct {
	ResponseType string `json:"response_type,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	Username string `json:"username,omitempty"`
	Code string `json:"code,omitempty"`
	State string `json:"state,omitempty"`
	Scope string `json:"scope,omitempty"`
	ConfirmScope bool `json:"confirm_scope,omitempty"`
	RedirectUri *string `json:"redirect_uri,omitempty"`
	AccessToken *AccessTokenResponseDTO `json:"access_token,omitempty"`
}