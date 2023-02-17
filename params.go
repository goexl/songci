package songci

import (
	"time"
)

type params struct {
	timeout time.Duration
	zinan   *algorithm
	basic   *algorithm

	id      string
	product string
	service string
}

func newParams() *params {
	return &params{
		timeout: 5 * time.Minute,
		zinan:   newAlgorithm(zinanName),
		basic:   newAlgorithm(basicName),

		product: product,
		service: service,
	}
}
