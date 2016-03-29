//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package auth

import (
	"lily"
)

const (
	DEFAULT_AUTH_COOKIE     = "LILYAUTH"
	DEFAULT_AUTH_LENGTH     = 10
	USER                    = "user"
)

var (
	cookieName   = DEFAULT_AUTH_COOKIE
	cookieLength = DEFAULT_AUTH_LENGTH
)

func init()  {
	lily.RegisterMiddleware("auth", Register)
}

func CheckAuth(request *lily.Request) {
	if sessionCookie, err := request.Cookie(cookieName); err == nil {
		request.Context[USER] = GetUserFromAuth(sessionCookie.Value)
	}
}

func SetSession(request *lily.Request, response *lily.Response) {
}

func Register(handler lily.IHandler) {
	handler.Initializer().Register(CheckAuth)
	handler.Finalizer().RegisterFinish(SetSession)
}