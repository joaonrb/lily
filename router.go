package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//
// Default router for Lily. Load the router in main app. Router must implement IRouter.
// Package router loads a string in format:
// "/path1/path2/:(?P<parameter>\d+)" ExampleController
//

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	urls = &route{paths: map[string]*route{}}
)

type route struct {
	next       *route
	paths      map[string]*route
	regex      *regexp.Regexp
	controller IController
}

func getController(uri []byte) (IController, map[string]interface{}) {
	if uri[0] == '/' {
		uri = uri[1:]
	}
	if len(uri) > 0 && uri[len(uri)-1] == '/' {
		uri = uri[:len(uri)-1]
	}
	params := map[string]interface{}{}
	way := urls
	for _, part := range bytes.Split(uri, []byte{'/'}) {
		value, exist := way.paths[string(part)]
		switch {
		case exist:
			way = value
		case way.regex != nil:
			match := way.regex.FindSubmatch(part)
			if len(match) > 0 {
				params[way.regex.SubexpNames()[1]] = string(match[0])
				way = way.next
			} else {
				return nil, nil
			}
		default:
			return nil, nil
		}
	}
	return way.controller, params
}

// Register an controller to a path
func Url(uri string, controller IController) error {
	controller.Init(controller)
	if uri[0] == '/' {
		uri = uri[1:]
	}
	if len(uri) > 0 && uri[len(uri)-1] == '/' {
		uri = uri[:len(uri)-1]
	}
	way := urls
	parts := strings.Split(uri, "/")
	var err error
	for _, part := range parts {
		if len(part) > 0 && part[0] == ':' {
			way.regex, err = regexp.Compile(part[1:])
			if err != nil {
				return err
			}
			if way.next == nil {
				way.next = &route{paths: map[string]*route{}}
			}
			way = way.next
		} else {
			if _, exist := way.paths[part]; !exist {
				way.paths[part] = &route{paths: map[string]*route{}}
			}
			way = way.paths[part]
		}
	}
	way.controller = controller
	return nil
}
