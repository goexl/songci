package songci

var _ = New

// New 创建新签名器
func New() *builder {
	return newBuilder()
}
