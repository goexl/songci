package songci

import (
	"time"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Timeout(timeout time.Duration) *builder {
	b.params.timeout = timeout

	return b
}
