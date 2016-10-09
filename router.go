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
    "strings"
)

var (
	urls = &route{}
)

type route struct {
	next       *route
	paths      map[string]IController
	regex      *regexp.Regexp
	controller IController
}

func getController(uri string) (IController, map[string]string) {
	if uri[0] == "/" { uri = uri[1:] }
	if uri[len(uri)-1] == "/" { uri = uri[:len(uri)-1] }
	params := map[string]string{}
	way := urls
	for part := range strings.Split(uri, "/") {
		if _, exist := way.paths[part]; !exist {
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

func Url(uri string, controller IController) error {
	if uri[0] == "/" { uri = uri[1:] }
	if uri[len(uri)-1] == "/" { uri = uri[:len(uri)-1] }
	way := urls
	parts := strings.Split(uri, "/")
	var err error
	for part := range parts {
		if part[0] == ":" {
			way.regex, err = regexp.Compile(part[1:])
			if err != nil {
				return err
			}
			if way.next == nil {
				way.next = &route{}
			}
			way = way.next
		} else {
			if _, exist := way.paths[part]; !exist {
				way.paths[part] = &route{}
			}
			way = way.paths[part]
		}
	}
	way.controller = controller
	return nil
}