//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"regexp"
	"strings"
)

type IRouter interface {
	Parse(string) (IController, map[string]string)
}

type IRoute interface {
	GetRoute(string, map[string]string) IRoute
	GetController() IController
}

func C()

type Router struct {
	head IRoute
}

func (self *Router) Parse(path string) (IController, map[string]string) {
	subpaths := strings.Split(path, "/")
	route := self.head
	params := map[string]string{}
	for _, subpath := range subpaths {
		route = route.GetRoute(subpath, params)
	}
	if route.GetController() != nil {
		return route.GetController(), params
	}
	return nil, nil
}

type Route struct {
	controller IController
}

func (self *Route) GetController() IController {
	return self.controller
}

type simpleRoute struct {
	Route
	routes map[string]IRoute
}

func (self *simpleRoute) GetRoute(route string, params map[string]string) IRoute {
	return self.routes[route]
}

type regexRoute struct {
	Route
	regex *regexp.Regexp
	route IRoute
}

func (self *regexRoute) GetRoute(route string, params map[string]string) IRoute {
	match := self.regex.FindStringSubmatch(route)
	if len(match) > 0 {
		params[self.regex.SubexpNames()[0]] = match[0]
		return self.route
	}
	return nil
}
