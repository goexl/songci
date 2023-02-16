package songci

var _ getter = (*stringGetter)(nil)

type stringGetter struct {
	credential string
}

func NewStringGetter(credential string) *stringGetter {
	return &stringGetter{
		credential: credential,
	}
}

func (sg *stringGetter) Get(_ string) (credential string, err error) {
	credential = sg.credential

	return
}
