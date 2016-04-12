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
	"github.com/go-macaron/cache"
	"strings"
	"fmt"
)

var (
	cacheEngine cache.Cache
)

func init() {
	var err error
	cacheEngine, err = cache.NewCacher("memory", cache.Options{})
	if err != nil {
		panic(err)
	}
}

const (
	DEFAULT_SESSION_COOKIE     = "LILYSESSION"
	DEFAULT_SESSION_LENGTH     = 10
	SESSION                    = "session"
	DEFAULT_SESSION_TIMEOUT    = 12  // In hours
)

var (
	cookieName    = DEFAULT_SESSION_COOKIE
	cookieLength  = DEFAULT_SESSION_LENGTH
	cookieTimeout = DEFAULT_SESSION_TIMEOUT
)

func init()  {
	lily.RegisterMiddleware("sessions", Register)
}

type session struct {
	Cookie  string
	session map[string]interface{}
}

func (self *session) Get(key string) interface{} {
	self.load()
	return self.session[key]
}

func (self *session) Set(key string, value interface{}) {
	self.load()
	self.session[key] = value
}

func (self *session) load() {
	if self.session == nil {
		if len(self.Cookie) != 0 {
			tmp := cacheEngine.Get(fmt.Sprintf("%s_%s", SESSION, self.Cookie))
			if tmp != nil {
				self.session = tmp.(map[string]interface{})
				return
			}
		}
		self.session = make(map[string]interface{})
	}
}

func GetSession(request *lily.Request) *session {
	return request.Context[SESSION].(*session)
}

func CheckSession(request *lily.Request) {
	request.Context[SESSION] = &session{string(request.Header.Cookie(cookieName)), nil}
}

func SetSession(request *lily.Request, response *lily.Response) {
	session := GetSession(request)
	if len(session.Cookie) == 0 {
		session.Cookie = lu.GenerateBase64String(cookieLength)
		cookie := &http.Cookie{
			Name: cookieName,
			Value: session.Cookie,
			Path: "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	}
	go cacheEngine.Put(session.Cookie, session.session, cookieTimeout * 60 * 60)
}

func Register(handler lily.IHandler) {
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
