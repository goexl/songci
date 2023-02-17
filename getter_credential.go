package songci

var _ getter = (*credentialGetter)(nil)

type credentialGetter struct {
	credential string
}

// NewCredential 普通凭据
func NewCredential(credential string) *credentialGetter {
	return &credentialGetter{
		credential: credential,
	}
}

func (cg *credentialGetter) Get(_ string, _ string) (credential string, err error) {
	credential = cg.credential

	return
}
