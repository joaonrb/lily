// Package lily
//
// Copyright (c) João Nuno. All rights reserved.
//
// Default router for Lily. Load the router in main app. Router must implement IRouter.
// Package router loads a string in format:
// `/path1/path2/:(?P<parameter>\d+)` ExampleController
//
package lily

import (
	"regexp"
	"strings"
)

var appRouter IRouter

func LoadRouter(path string) {

	appRouter = router
}

type IRouter interface {
	Parse(string) (IController, map[string]string)
	GetPath(string, map[string]string)
}

type IRoute interface {
	GetRoute(string, map[string]string) IRoute
	GetController() IController
}


type router struct {
	head IRoute
}

func (self *router) Parse(path string) (IController, map[string]string) {
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

type route struct {
	controller IController
}

func (self *route) GetController() IController {
	return self.controller
}

type simpleRoute struct {
	route
	routes map[string]IRoute
}

func (self *simpleRoute) GetRoute(route string, params map[string]string) IRoute {
	return self.routes[route]
}

type regexRoute struct {
	route
	regex *regexp.Regexp
	path IRoute
}

func (self *regexRoute) GetRoute(route string, params map[string]string) IRoute {
	match := self.regex.FindStringSubmatch(route)
	if len(match) > 0 {
		params[self.regex.SubexpNames()[0]] = match[0]
		return self.path
	}
	return nil
}