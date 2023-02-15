package songci

type algorithm interface {
	name() string
	signature() string
	headers() headers
}
