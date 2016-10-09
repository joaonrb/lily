//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
)

func core(ctx *fasthttp.RequestCtx) {
	defer func() {
		if recovery := recover(); recovery != nil {
			status := fasthttp.StatusInternalServerError
			ctx.Response.SetStatusCode(status, fasthttp.StatusMessage(status))
		}
	}()
	controller, args := getController(ctx.Path())
	controller.Handle(controller, ctx, args)
}