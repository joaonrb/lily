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
	DEFAULT_SESSION_COOKIE     = "LILYGUEST"
	DEFAULT_SESSION_AGE        = time.Hour
	DEFAULT_SESSION_LENGTH = 11
	SESSION                    = "session"
	SET_SESSION                = "set-session"
	COOKIE                     = "cookie"
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
		session, exist := sessionStore.GetSession(sessionCookie.Value)
		if !exist {
			session = Session{COOKIE: lu.GenerateBase64String(cookieLength)}
			request.Context[SET_SESSION] = true
		}
		request.Context[SESSION] = session
	}
}

func SetSession(request *lily.Request, response *lily.Response) {
	session := GetSession(request)
	if session[SESSION].(bool) {
		sessionStore.SetSession(session[COOKIE].(string), time.Now().UTC().Add(maxAge), session)
		cookie := &http.Cookie{
			Name: cookieName,
			Value: session[COOKIE].(string),
			MaxAge: maxAge.Seconds(),
			Path: "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	} else {
		sessionStore.UpdateSession(session[COOKIE].(string), session)
	}
}

func Register(handler lily.IHandler) {
	LoadSessionStore()
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
