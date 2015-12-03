//
// Copyright (c) João Nuno. All rights reserved.
//
package lily


type RequestMiddleware func(*Request)

type ResponseMiddleware func(*Request, *Response)
