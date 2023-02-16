package songci

type getter interface {
	// Get 获取凭据
	Get(id string) (credential string, err error)
}
