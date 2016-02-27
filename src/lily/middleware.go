//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily


type RequestMiddleware func(*Request)

type ResponseMiddleware func(*Request, *Response)
