package songci

type coreParams struct {
	host    string
	method  string
	uri     string
	headers headers
	payload []byte
}

func newCoreParams() *coreParams {
	return new(coreParams)
}
