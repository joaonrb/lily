//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

var resgistedMiddleware = map[string]func(handler IHandler){}

func RegisterMiddleware(name string, middleware func(handler IHandler)) {
	resgistedMiddleware[name] = middleware
}

type RequestMiddleware func(*Request)

type ResponseMiddleware func(*Request, *Response)
