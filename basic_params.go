package songci

type basicParams struct {
	password string
}

func newBasicParams() *basicParams {
	return new(basicParams)
}
