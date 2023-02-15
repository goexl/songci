package songci

// Verifier 验证器
type Verifier interface {
	// Verify 验证
	Verify() (verified bool, err error)
}
