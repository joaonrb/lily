//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package sessions

import (
	"lily"
	lu "lily/utils"
	"time"
	"net/http"
)

const (
	DEFAULT_SESSION_COOKIE     = "LILYSESSION"
	DEFAULT_SESSION_AGE        = time.Hour
	DEFAULT_SESSION_LENGTH     = 10
	SESSION                    = "session"
)

var (
	cookieName   = DEFAULT_SESSION_COOKIE
	cookieLength = DEFAULT_SESSION_LENGTH
	maxAge       = DEFAULT_SESSION_AGE
)

func init()  {
	lily.RegisterMiddleware("sessions", Register)
}

type Session map[string] interface{}

func GetSession(request *lily.Request) Session {
	return request.Context[SESSION].(Session)
}

func CheckSession(request *lily.Request) {
	if sessionCookie, err := request.Cookie(cookieName); err == nil && time.Now().UTC().Before(sessionCookie.Expires) {
		request.Context[SESSION] = sessionCookie.Value
	}
}

func SetSession(request *lily.Request, response *lily.Response) {
	session := GetSession(request)
	if _, exist := session[SESSION]; !exist {
		cookie := &http.Cookie{
			Name: cookieName,
			Value: lu.GenerateBase64String(cookieLength),
			MaxAge: maxAge.Seconds(),
			Path: "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	}
}

func Register(handler lily.IHandler) {
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
