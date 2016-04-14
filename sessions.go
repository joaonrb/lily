//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
// Check cache documentation here: https://go-macaron.com/docs/middlewares/cache
//
package lily

import (
	"net/http"
	"github.com/go-macaron/cache"
	"fmt"
)

var (
	cacheEngine cache.Cache
)

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
	var err error
	cacheEngine, err = cache.NewCacher("memory", cache.Options{})
	if err != nil {
		panic(err)
	}
	RegisterMiddleware("sessions", Register)
}

func LoadCache(conf *Settings) {
	cacheConf := conf.Apps["cache"].(map[interface{}]interface{})
	var options cache.Options
	var err error
	adapterConfig, exist := cacheConf["adapter_config"]
	if exist {
		options = cache.Options{AdapterConfig: adapterConfig.(string)}
	} else {
		options = cache.Options{}
	}
	cacheEngine, err = cache.NewCacher(cacheConf["type"].(string), options)
	if err != nil {
		panic(err)
	}
}

type session struct {
	Cookie  string
	session map[string]interface{}
	get     func(*session, string) interface{}
	set     func(*session, string, interface{})
}

func NewSession(cookie string) *session {
	return &session{
		Cookie: cookie,
		session: nil,
		get: loadGet,
		set: loadSet,
	}
}

func _get(self *session, key string) interface{} {
	return self.session[key]
}

func (self *session) Get(key string) interface{} {
	return self.get(self, key)
}

func _set(self *session, key string, value interface{}) {
	self.session[key] = value
}

func (self *session) Set(key string, value interface{}) {
	self.set(self, key, value)
}

func (self *session) Load() {
	defer func() {
		self.get, self.set = _get, _set
	}()
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

func loadGet(self *session, key string) interface{} {
	self.Load()
	return self.Get(key)
}

func loadSet(self *session, key string, value interface{}) {
	self.Load()
	self.Set(key, value)
}

func GetSession(request *Request) *session {
	return request.Context[SESSION].(*session)
}

func CheckSession(request *Request) {
	request.Context[SESSION] = NewSession(string(request.Header.Cookie(cookieName)))
}

func SetSession(request *Request, response *Response) {
	session := GetSession(request)
	if len(session.Cookie) == 0 {
		session.Cookie = GenerateBase64String(cookieLength)
		cookie := &http.Cookie{
			Name: cookieName,
			Value: session.Cookie,
			Path: "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	}
	go cacheEngine.Put(session.Cookie, session.session, int64(cookieTimeout * 60 * 60))
}

func Register(handler IHandler) {
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
