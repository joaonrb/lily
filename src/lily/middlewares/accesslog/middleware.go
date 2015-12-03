//
// Copyright (c) João Nuno. All rights reserved.
//
package accesslog

import (
	"lily"
	"time"
	"strconv"
	"lily/middlewares/auth"
)

const (
	REQUEST_START = "__start__"
	TIME_FORMAT = "02/Jan/2006:15:04:05Z0700"
)

func InitRequestForLog(request *lily.Request) {
	request.Context[REQUEST_START] = time.Now().UTC()
}

func FinishRequestForLog(request *lily.Request, response *lily.Response) {
	request.Context[REQUEST_START] = time.Now().UTC()

	status := response.Status
	if status == 0 {
		status = 404
	}
	bodyLen, _ := strconv.ParseInt(response.RW.Header().Get("Content-Length"), 10, 64)
	user := auth.GetUser(request)
	ip := request.RemoteAddr
	method := request.Method
	path := request.RequestURI
	httpVersion := request.Proto
	start := request.Context[REQUEST_START].(time.Time)
	log.Info(
		"%s %s [%s] \"%s %s %s\" %d %d %dms", ip, user, time.Now().Format(TIME_FORMAT), method, path, httpVersion,
		status, bodyLen, time.Since(start).Nanoseconds() / 1000000,
	)
}

func Register(handler lily.IHandler) {
	LoadAccessLogger()
	handler.Initializer().Register(InitRequestForLog)
	handler.Finalizer().RegisterFinish(FinishRequestForLog)
}