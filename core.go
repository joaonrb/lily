package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"github.com/valyala/fasthttp"
)

// The core handler to be called.
func CoreHandler(ctx *fasthttp.RequestCtx) {
	controller, args := getController(ctx.Path())
	if controller == nil {
		response := Http404()
		sendResponse(ctx, response)
	} else {
		controller.Handle(ctx, args)
	}
}
