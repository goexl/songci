package songci

import (
	"time"
)

type params struct {
	timeout time.Duration
	zinan   *algorithm
	basic   *algorithm
}

func newParams() *params {
	return &params{
		timeout: 5 * time.Minute,
		zinan:   newAlgorithm(zinanName),
		basic:   newAlgorithm(basicName),
	}
}
