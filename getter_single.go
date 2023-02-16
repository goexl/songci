package songci

var _ getter = (*singleGetter)(nil)

type singleGetter struct {
	credential string
}

// NewSingleGetter 单一凭据获取
func NewSingleGetter(credential string) *singleGetter {
	return &singleGetter{
		credential: credential,
	}
}

func (sg *singleGetter) Get(_ string) (credential string, err error) {
	credential = sg.credential

	return
}
