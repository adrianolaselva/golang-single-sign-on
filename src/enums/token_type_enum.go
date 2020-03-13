package enums

type TokenTypeEnum int

const (
	Bearer TokenTypeEnum = iota
)

func (g TokenTypeEnum) String() string {
	toString := map[TokenTypeEnum]string{
		Bearer:  "Bearer",
	}
	return toString[g]
}