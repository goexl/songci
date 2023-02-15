package songci

import (
	"time"
)

type builder struct {
	// 超时
	timeout time.Duration
}

func newBuilder(timeout time.Duration) *builder {
	return &builder{
		timeout: timeout,
	}
}

func (b *builder) Timeout(timeout time.Duration) *builder {
	b.timeout = timeout

	return b
}
