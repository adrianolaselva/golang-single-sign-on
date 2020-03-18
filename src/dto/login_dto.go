package dto

type LoginDTO struct {
	ResponseType string `json:"response_type,omitempty" schema:"response_type,omitempty"`
	ClientID string `json:"client_id,omitempty" schema:"client_id,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty" schema:"redirect_uri,omitempty"`
	Scope string `json:"scope,omitempty" schema:"scope,omitempty"`
	State string `json:"state,omitempty" schema:"state,omitempty"`
	Username string `json:"username,omitempty" schema:"username,omitempty"`
	Password string `json:"password,omitempty" schema:"password,omitempty"`
}