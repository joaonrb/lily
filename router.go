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
	"bytes"
)

var (
	urls = &route{}
)

type route struct {
	next       *route
	paths      map[string]*route
	regex      *regexp.Regexp
	controller IController
}

func getController(uri []byte) (IController, map[string]string) {
	if uri[0] == '/' { uri = uri[1:] }
	if uri[len(uri)-1] == '/' { uri = uri[:len(uri)-1] }
	params := map[string]string{}
	way := urls
	for _, part := range bytes.Split(uri, []byte{'/'}) {
		if _, exist := way.paths[string(part)]; !exist {
			match := way.regex.FindSubmatch(part)
			if len(match) > 0 {
				params[way.regex.SubexpNames()[1]] = string(match[0])
				way = way.next
			} else {
				return nil, nil
			}
		} else {
			way = way.paths[string(part)]
		}
	}
	if way != nil {
		return way.controller, params
	}
	return nil, nil
}

func Url(uri string, controller IController) error {
	if uri[0] == '/' { uri = uri[1:] }
	if uri[len(uri)-1] == '/' { uri = uri[:len(uri)-1] }
	way := urls
	parts := strings.Split(uri, "/")
	var err error
	for _, part := range parts {
		if part[0] == ':' {
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