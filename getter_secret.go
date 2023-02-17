package songci

var _ getter = (*secretGetter)(nil)

type secretGetter struct {
	credential string
}

// NewSecret 普通密钥
func NewSecret(secret string) *secretGetter {
	return &secretGetter{
		credential: secret,
	}
}

func (sg *secretGetter) Get(_ string, _ string) (credential string, err error) {
	credential = sg.credential

	return
}
