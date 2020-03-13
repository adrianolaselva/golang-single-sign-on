package enums

type GrantTypeEnum string

const (
	GrantTypeAuthorizationCode GrantTypeEnum = "authorization_code"
	GrantTypeClientCredentials GrantTypeEnum = "client_credentials"
	GrantTypePassword GrantTypeEnum = "password"
	GrantTypeImplicit GrantTypeEnum = "implicit"
	GrantTypeRefreshToken GrantTypeEnum = "refresh_token"
)

func (g GrantTypeEnum) Parse(grantType string) GrantTypeEnum {
	return GrantTypeEnum(grantType)
}

func (g GrantTypeEnum) String() string {
	toString := map[GrantTypeEnum]string{
		GrantTypeAuthorizationCode:  "authorization_code",
		GrantTypeClientCredentials:  "client_credentials",
		GrantTypePassword: "password",
		GrantTypeImplicit: "implicit",
		GrantTypeRefreshToken: "refresh_token",
	}
	return toString[g]
}