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

func (b *builder) Zinan() *algorithmBuilder {
	return newAlgorithmBuilder(b, b.params.zinan)
}

func (b *builder) Basic() *algorithmBuilder {
	return newAlgorithmBuilder(b, b.params.basic)
}

func (b *builder) Build() *Songci {
	return newSongci(b.params)
}
