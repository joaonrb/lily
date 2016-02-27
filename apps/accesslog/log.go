//
// Author João Nuno.
//
package accesslog

import (
	"github.com/op/go-logging"
	"lily"
	"fmt"
	"os"
	"io"
)

var log *logging.Logger

func LoadAccessLogger() {
	settings := lily.Configuration
	if settings.AccessLog.Type == "" {
		panic(fmt.Errorf("Trying to use accesslog middleware without configure the path for it"))
	}
	log = logging.MustGetLogger("accesslog")
	var out io.Writer
	switch settings.AccessLog.Type {
	case lily.CONSOLE:
		out = os.Stdout
	case lily.FILE:
		out = lily.OpenRotatorFile(settings.AccessLog.Path)
		if fmt.Printf("# Access log recording at %s", settings.AccessLog.Path)
	}
	logger := logging.NewLogBackend(out, "", 0)
	beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter("%{message}"))
	beLevel := logging.AddModuleLevel(beFormatter)
	beLevel.SetLevel(logging.INFO, "")
	log.SetBackend(beLevel)
}