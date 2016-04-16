package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

var middlewares = map[string]func(handler IHandler){}

func RegisterMiddleware(name string, middleware func(handler IHandler)) {
	middlewares[name] = middleware
}

type RequestMiddleware func(*Request)

type ResponseMiddleware func(*Request, *Response)
