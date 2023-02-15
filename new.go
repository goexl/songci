package songci

import (
	"time"
)

var _ = New

func New() *builder {
	return newBuilder(5 * time.Minute)
}
