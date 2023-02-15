package songci

var _ Verifier = (*zinanVerifier)(nil)

type zinanVerifier struct {
	params *zinanParams
}

func newZinanVerifier(params *zinanParams) *zinanVerifier {
	return &zinanVerifier{
		params: params,
	}
}

func (zv *zinanVerifier) Verify() (verified bool, err error) {
	return
}
