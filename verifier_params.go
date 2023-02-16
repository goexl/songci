package songci

type verifierParams struct {
	method  string
	uri     string
	query   string
	headers headers
	payload []byte
}

func newVerifierParams() *verifierParams {
	return new(verifierParams)
}
