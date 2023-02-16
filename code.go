package songci

const (
	codeSignatureError     uint8 = 1
	codeNotImplement       uint8 = 2
	codeGetCredentialError uint8 = 3

	codeNoContentTypeHeader uint8 = 10
	codeNoHostHeader        uint8 = 11

	codeTimestampFormatError uint8 = 20
	codeTimeout              uint8 = 21
)
