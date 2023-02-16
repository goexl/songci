package songci

type algorithm struct {
	name string
}

func newAlgorithm(name string) *algorithm {
	return &algorithm{
		name: name,
	}
}
