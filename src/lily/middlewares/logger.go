//
// Copyright (c) João Nuno. All rights reserved.
//
package middlewares

import (
	"lily"
	"time"
)

const (
	REQUEST_START = "__start__"
)

func InitRequestForLog(request *lily.Request) {
	request.Context[REQUEST_START] = time.Now().UTC()
}