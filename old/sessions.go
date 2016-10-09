//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily


import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"time"
	"net/http"
)

// Fazer middleware para header e para cookie

const (
	DEFAULT_SESSION_HEADER = "x-session"
	DEFAULT_SESSION_COOKIE = "session"
)

var (
	SESSION_HEADER = DEFAULT_SESSION_HEADER
	SESSION_COOKIE = DEFAULT_SESSION_COOKIE
)


func InitSession(request *Request) {
	if header := request.Header.Peek(SESSION_HEADER)
	request.Context[RequestStart] = time.Now().UTC()
}

func FinishRequestForLog(request *Request, response *Response) {
	request.Context[RequestStart] = time.Now().UTC()

	status := response.Status
	if status == 0 {
		status = http.StatusNotFound
	}
	bodyLen := len(response.Body)
	ip := request.RemoteAddr()
	method := request.Method()
	path := string(request.RequestURI())
	httpVersion := "HTTP1.1"
	if !request.Header.IsHTTP11() {
		httpVersion = "HTTP1.0'"
	}
	start := request.Context[RequestStart].(time.Time)
	accessLog.Infof(
		"%s [%s] \"%s %s %s\" %d %d %s", ip, time.Now().Format(TimeFormat), method, path, httpVersion,
		status, bodyLen, time.Since(start).String(),
	)
}

func main() {
	key := []byte("example key 1234")
	ciphertext, _ := hex.DecodeString("22277966616d9bc47177bd02603d08c9a67d5380d0fe8cf3b44438dff7b9")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)
}