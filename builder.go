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
