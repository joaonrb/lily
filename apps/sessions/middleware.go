//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package sessions

import (
	"lily"
	lu "lily/utils"
	"net/http"
)

const (
	DEFAULT_SESSION_COOKIE     = "LILYSESSION"
	DEFAULT_SESSION_LENGTH     = 10
	SESSION                    = "session"
)

var (
	cookieName   = DEFAULT_SESSION_COOKIE
	cookieLength = DEFAULT_SESSION_LENGTH
)

func init()  {
	lily.RegisterMiddleware("sessions", Register)
}

func GetSession(request *lily.Request) string {
	return request.Context[SESSION].(string)
}

func CheckSession(request *lily.Request) {
	request.Context[SESSION] = request.Header.Cookie(cookieName)
}

func SetSession(request *lily.Request, response *lily.Response) {
	session := GetSession(request)
	if _, exist := session[SESSION]; !exist {
		cookie := &http.Cookie{
			Name: cookieName,
			Value: lu.GenerateBase64String(cookieLength),
			Path: "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	}
}

func Register(handler lily.IHandler) {
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
