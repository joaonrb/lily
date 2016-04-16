package lily

//
// Author João Nuno.
//

import (
	"fmt"
	"github.com/op/go-logging"
	"io"
	"net/http"
	"os"
	"time"
)

var accessLog *logging.Logger

func LoadAccessLogger() {
	settings := Configuration
	if settings.AccessLog.Type == "" {
		panic(fmt.Errorf("Trying to use accesslog middleware without configure the path for it"))
	}
	accessLog = logging.MustGetLogger("accesslog")
	var out io.Writer
	switch settings.AccessLog.Type {
	case CONSOLE:
		out = os.Stdout
	case FILE:
		out = OpenRotatorFile(settings.AccessLog.Path)
		fmt.Printf("# Access log recording at %s", settings.AccessLog.Path)
	}
	logger := logging.NewLogBackend(out, "", 0)
	beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter("%{message}"))
	beLevel := logging.AddModuleLevel(beFormatter)
	beLevel.SetLevel(logging.INFO, "")
	accessLog.SetBackend(beLevel)
}

const (
	RequestStart = "__start__"
	TimeFormat   = "02/Jan/2006:15:04:05Z0700"
)

func init() {
	RegisterMiddleware("accesslog", AccessLogRegister)
}

func InitRequestForLog(request *Request) {
	request.Context[RequestStart] = time.Now().UTC()
}

func FinishRequestForLog(request *Request, response *Response) {
	request.Context[RequestStart] = time.Now().UTC()

	status := response.Status
	if status == 0 {
		status = http.StatusNotFound
	}
	bodyLen := len(response.Body)
	ip := request.RemoteAddr()
	method := request.Method()
	path := string(request.RequestURI())
	httpVersion := "HTTP1.1"
	if !request.Header.IsHTTP11() {
		httpVersion = "HTTP1.0'"
	}
	start := request.Context[RequestStart].(time.Time)
	accessLog.Infof(
		"%s [%s] \"%s %s %s\" %d %d %s", ip, time.Now().Format(TimeFormat), method, path, httpVersion,
		status, bodyLen, time.Since(start).String(),
	)
}

func AccessLogRegister(handler IHandler) {
	LoadAccessLogger()
	handler.Initializer().Register(InitRequestForLog)
	handler.Finalizer().RegisterFinish(FinishRequestForLog)
}
