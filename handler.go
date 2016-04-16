package lily
//
// Author João Nuno.
//

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

var mainHandler IHandler

func ForceRegisterHandler(handler IHandler) {
	mainHandler = handler
}

func RegisterHandler(handler IHandler) bool {
	if mainHandler != nil {
		ForceRegisterHandler(handler)
		return true
	}
	return false
}

func defaultHandler() IHandler {
	return NewHandler(NewRequestInitializer(), NewFinalizer())
}

type IHandler interface {
	ServeHTTP(context *fasthttp.RequestCtx)
	RegisterStaticPath(uri, path string)
	Initializer() IInitializer
	Finalizer() IFinalizer
}

/**
 * Implements http.Handler
 */
type Handler struct {
	init   IInitializer
	finish IFinalizer
	static map[string]fasthttp.RequestHandler
}

func NewHandler(init IInitializer, finish IFinalizer) *Handler {
	return &Handler{init, finish, map[string]fasthttp.RequestHandler{}}
}

func (self *Handler) ServeHTTP(context *fasthttp.RequestCtx) {
	var response *Response
	var lilyRequest *Request
	defer func() {
		err := recover()
		if err != nil {
			switch err := err.(type) {
			case IHttpError:
				response = err.ToResponse()
				Debug("Http error caught with status %d and message %s", response.Status, response.Body)
			case error:
				response = NewHttp500(err.Error()).ToResponse()
				Error(err.Error())
			default:
				msg := fmt.Sprintf("%x", err)
				response = NewHttp500(msg).ToResponse()
				Error(msg)
			}
			response.FastHttpResponse = &(context.Response)
		}
		context.SetContentType(lilyRequest.Context[contentType].(string))
		self.finish.Finish(lilyRequest, response)
	}()
	lilyRequest = self.init.Start(context)

	path := string(lilyRequest.URI().Path())
	if static, exist := self.static[path]; exist {
		static(context)
		return
	}
	controller, params, err := mainRouter.Parse(path)
	if err != nil {
		panic(err)
	}
	for _, middleware := range controller.PreMiddleware() {
		middleware(lilyRequest)
	}
	response = controller.Handle(controller, lilyRequest, params)
	response.FastHttpResponse = &(context.Response)
	for _, middleware := range controller.PosMiddleware() {
		middleware(lilyRequest, response)
	}
}

func (self *Handler) RegisterStaticPath(uri, path string) {
	self.static[uri] = (&fasthttp.FS{
		// Path to directory to serve.
		Root: path,

		// Generate index pages if client requests directory contents.
		// GenerateIndexPages: true,

		// Enable transparent compression to save network traffic.
		Compress: true,
	}).NewRequestHandler()
}

func (self *Handler) Initializer() IInitializer { return self.init }

func (self *Handler) Finalizer() IFinalizer { return self.finish }
