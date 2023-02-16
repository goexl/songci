package songci

type authorizer interface {
	sign() (signature string, err error)

	authorize() (authorize string, err error)
}
