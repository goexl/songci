package songci

type getter interface {
	// Get 获取凭据
	Get(scheme string, id string) (credential string, err error)
}
