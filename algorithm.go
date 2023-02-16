package songci

type algorithm struct {
	scheme string
}

func newAlgorithm(scheme string) *algorithm {
	return &algorithm{
		scheme: scheme,
	}
}
