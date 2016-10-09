//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
// Default router for Lily. Load the router in main app. Router must implement IRouter.
// Package router loads a string in format:
// "/path1/path2/:(?P<parameter>\d+)" ExampleController
//
package lily

import (
	"regexp"
)

var (
	urls = map[string]route{}
)

type route struct {
	next       route
	paths      map[string]IController
	regex      *regexp.Regexp
	controller IController
}

func getController(uri string) (IController, map[string]string) {
	var (
		parts  []string
		way    *route
	    exist  bool
		params map[string]string
	)
	parts = uri.Split("/")
	way, exist = urls[parts[0]]
	if !exist {
		return
	}
	for part := range parts[1:] {
		if _, exist = way.paths[part]; !exist {
			match := way.regex.FindStringSubmatch(part)
			if len(match) > 0 {
				params[way.regex.SubexpNames()[1]] = match[0]
				way = way.next
			} else {
				return
			}
		} else {
			way = way.paths[part]
		}
	}
	if way != nil {
		return way.controller, params
	}
	return
}