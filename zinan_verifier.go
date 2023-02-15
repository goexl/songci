package songci

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var _ Verifier = (*zinanVerifier)(nil)

type zinanVerifier struct {
	params *zinanParams
}

func newZinanVerifier(params *zinanParams) *zinanVerifier {
	return &zinanVerifier{
		params: params,
	}
}

func (zv *zinanVerifier) Verify() (verified bool, err error) {
	// 第一步，验证参数
	if codes, verified := zv.params.validate(); !verified {
		err = exc.NewField("参数不合法", field.New("codes", codes))
	}
	if nil != err {
		return
	}

	// 第二步，组装签名字符串

	return
}

func (zv *zinanVerifier) request() (request string) {
	str := new(strings.Builder)
	str.WriteString(zv.params.method)
	str.WriteString("\n")

	return
}
func (r *Rna) setHeader(request *resty.Request, host string, action core.Action, payload []byte) {
	algorithm := "TC3-HMAC-SHA256"
	service := "fiv"
	version := "2019-05-21"
	region := "ap-guangzhou"
	timestamp := time.Now().Unix()
	method := "POST"
	uri := "/"
	query := ""
	canonicalHeaders := "content-type:application/json; charset=utf-8\n" + "host:" + host + "\n"
	signedHeaders := "content-type;host"

	hashedRequestPayload := r.sha256Hex(string(payload))
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		method,
		uri,
		query,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload,
	)

	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := r.sha256Hex(canonicalRequest)
	stringSign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest,
	)

	secretDate := r.hmacSha256(date, "TC3"+r.tencent.Secret)
	secretService := r.hmacSha256(service, secretDate)
	secretSigning := r.hmacSha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(r.hmacSha256(stringSign, secretSigning)))
	authorization := fmt.Sprintf("Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		r.tencent.Id,
		credentialScope,
		signedHeaders,
		signature,
	)

	// 不能直通设置头，会被Resty中间件覆盖
	request.SetAuthScheme(algorithm)
	request.SetAuthToken(authorization)
	request.SetHeader("Content-Type", "application/json; charset=utf-8")
	request.SetHeader("Host", host)
	request.SetHeader("X-TC-Action", string(action))
	request.SetHeader("X-TC-Timestamp", strconv.FormatInt(timestamp, 10))
	request.SetHeader("X-TC-Version", version)
	request.SetHeader("X-TC-Region", region)
}

func (r *Rna) sha256Hex(from string) (to string) {
	bytes := sha256.Sum256([]byte(from))
	to = hex.EncodeToString(bytes[:])

	return
}

func (r *Rna) hmacSha256(s, key string) (to string) {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	to = string(hashed.Sum(nil))

	return
}
