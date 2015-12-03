//
// Copyright (c) João Nuno. All rights reserved.
//
package accesslog

import (
	"github.com/op/go-logging"
	"lily"
	"fmt"
)

var log *logging.Logger

func LoadAccessLogger() {
	settings := lily.Settings
	if settings.AccessLog.Path == "" {
		panic(fmt.Errorf("Trying to use accesslog middleware without configure the path for it"))
	}
	log = logging.MustGetLogger("accesslog")
	out := lily.OpenRotatorFile(settings.AccessLog.Path)
	logger := logging.NewLogBackend(out, "", 0)
	beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter("%{message}"))
	beLevel := logging.AddModuleLevel(beFormatter)
	beLevel.SetLevel(logging.INFO, "")
	log.SetBackend(beLevel)
}