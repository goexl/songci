package songci

type codeParams struct {
	scheme string
	typ    codeType
}

func newCodeParams() *codeParams {
	return &codeParams{
		scheme: https,
		typ:    codeTypeCurl,
	}
}
