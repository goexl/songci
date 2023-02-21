package songci

type authorizer interface {
	scheme() string

	resolve(authorization string) (codes codes)

	sign() (signature string, codes codes)

	token() (token string, codes codes)

	check(signature string) bool
}
