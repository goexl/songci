package songci

type authorizer interface {
	scheme() string

	resolve(authorization string) (codes []uint8)

	sign() (signature string, codes []uint8)

	token() (token string, codes []uint8)

	check(signature string) bool
}
