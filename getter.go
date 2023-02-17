package songci

type getter interface {
	// Get 获取凭据
	Get(scheme string, id string) (secret string, err error)
}
