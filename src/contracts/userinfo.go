package contracts

type UserInfo struct {
	Sub string `json:"sub,omitempty"`
	Name string `json:"name,omitempty"`
	Iat int32 `json:"iat,omitempty"`
	Email string `json:"email,omitempty"`
}
