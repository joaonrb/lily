package lily
//
// Author João Nuno.
//
// joaonrb@gmail.com
//
// Check cache documentation here: https://go-macaron.com/docs/middlewares/cache
//

import (
	"fmt"
	"github.com/go-macaron/cache"
	"net/http"
)

var (
	cacheEngine cache.Cache
)

const (
	defaultSessionCookie  = "LILYSESSION"
	defaultSessionLength  = 10
	sessionKey            = "session"
	defaultSessionTimeout = 12 // In hours
)

var (
	cookieName    = defaultSessionCookie
	cookieLength  = defaultSessionLength
	cookieTimeout = defaultSessionTimeout
)

func init() {
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
		Cookie:  cookie,
		session: nil,
		get:     loadGet,
		set:     loadSet,
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
			tmp := cacheEngine.Get(fmt.Sprintf("%s_%s", sessionKey, self.Cookie))
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
	return request.Context[sessionKey].(*session)
}

func CheckSession(request *Request) {
	request.Context[sessionKey] = NewSession(string(request.Header.Cookie(cookieName)))
}

func SetSession(request *Request, response *Response) {
	session := GetSession(request)
	if len(session.Cookie) == 0 {
		session.Cookie = GenerateBase64String(cookieLength)
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: session.Cookie,
			Path:  "/",
		}
		response.Headers["Set-Cookie"] = cookie.String()
	}
	go cacheEngine.Put(session.Cookie, session.session, int64(cookieTimeout*60*60))
}

func Register(handler IHandler) {
	handler.Initializer().Register(CheckSession)
	handler.Finalizer().RegisterFinish(SetSession)
}
