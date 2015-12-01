//
// Copyright (c) Telefonica I+D. All rights reserved.
//
package lily

import (
	"net/http"
)


// turn this into object

func LilyHandler(responseWriter http.ResponseWriter, request *http.Request) {

	lilyRequest := startRequest(request)  // Request-seption

	route := processPath(lilyRequest)
}