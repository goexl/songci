package songci

type verifierParams struct {
	credential string
	method     string
	uri        string
	query      string
	headers    headers
	payload    []byte
}

func newVerifierParams(credential string) *verifierParams {
	return &verifierParams{
		credential: credential,
	}
}
