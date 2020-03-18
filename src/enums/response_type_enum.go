package enums

type ResponseTypeEnum string

const (
	ResponseTypeAuthorizationCode ResponseTypeEnum = "code"
	ResponseTypeToken ResponseTypeEnum = "token"
)

func (g ResponseTypeEnum) Parse(responseType string) ResponseTypeEnum {
	return ResponseTypeEnum(responseType)
}

func (g ResponseTypeEnum) String() string {
	toString := map[ResponseTypeEnum]string{
		ResponseTypeAuthorizationCode:  "code",
		ResponseTypeToken: "token",
	}
	return toString[g]
}