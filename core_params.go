package songci

type coreParams struct {
	host    string
	method  string
	url     string
	query   string
	headers headers
	payload []byte
}

func newCoreParams() *coreParams {
	return new(coreParams)
}
