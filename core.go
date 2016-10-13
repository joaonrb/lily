// Package lily
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
)

func CoreHandler(ctx *fasthttp.RequestCtx) {
	controller, args := getController(ctx.Path())
	if controller == nil {
		response := Http404()
		sendResponse(ctx, response)
	} else {
		controller.Handle(ctx, args)
	}
}
