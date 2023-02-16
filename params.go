package songci

import (
	"time"
)

type params struct {
	// 超时
	timeout time.Duration
}

func newParams() *params {
	return &params{
		timeout: 5 * time.Minute,
	}
}
