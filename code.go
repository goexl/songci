package songci

const (
	codeSignatureError uint8 = 1
	codeNotImplement         = 2

	codeNoContentTypeHeader uint8 = 10
	codeNoHostHeader        uint8 = 11

	codeTimestampFormatError uint8 = 20
	codeTimeout              uint8 = 21
)
