package songci

type authorizer interface {
	scheme() string

	unzip(token string) (codes []uint8)

	sign() (signature string, codes []uint8)

	token() (token string, codes []uint8)

	signature() string
}
