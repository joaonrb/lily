//
// Author João Nuno.
//
package lily

import (
	"github.com/op/go-logging"
	"fmt"
	"os"
	"io"
	"time"
	"net/http"
)

var accesslog *logging.Logger

func LoadAccessLogger() {
	settings := Configuration
	if settings.AccessLog.Type == "" {
		panic(fmt.Errorf("Trying to use accesslog middleware without configure the path for it"))
	}
	accesslog = logging.MustGetLogger("accesslog")
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
	accesslog.SetBackend(beLevel)
}

const (
	REQUEST_START = "__start__"
	TIME_FORMAT = "02/Jan/2006:15:04:05Z0700"
)

func init()  {
	RegisterMiddleware("accesslog", AccesslogRegister)
}

func InitRequestForLog(request *Request) {
	request.Context[REQUEST_START] = time.Now().UTC()
}

func FinishRequestForLog(request *Request, response *Response) {
	request.Context[REQUEST_START] = time.Now().UTC()

	status := response.Status
	if status == 0 {
		status = http.StatusNotFound
	}
	bodyLen := len(response.Body)
	ip := request.RemoteAddr()
	method := request.Method()
	path := string(request.RequestURI())
	httpVersion := "HTTP1.1"; if !request.Header.IsHTTP11() { httpVersion = "HTTP1.0'" }
	start := request.Context[REQUEST_START].(time.Time)
	accesslog.Infof(
		"%s [%s] \"%s %s %s\" %d %d %s", ip, time.Now().Format(TIME_FORMAT), method, path, httpVersion,
		status, bodyLen, time.Since(start).String(),
	)
}

func AccesslogRegister(handler IHandler) {
	LoadAccessLogger()
	handler.Initializer().Register(InitRequestForLog)
	handler.Finalizer().RegisterFinish(FinishRequestForLog)
}