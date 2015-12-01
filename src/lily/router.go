//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"regexp"
	"strings"
	"fmt"
)

type Route struct {
	route      map[string]*Route
	regex      regexp.Regexp
	regexRoute *Route
	controller IController
}

func NewRoute() *Route{
	return &Route{map[string]Route{}, "", nil, nil}
}

func (self *Route) Controller(controller IController) {
	self.controller = controller
}

func (self *Route) C(controller IController) {
	self.Controller(controller)
}

func (self *Route) ParameterRoute(subpath string, route Route) {
	self.regex = regexp.MustCompile(subpath)
	self.regexRoute = route
}

func (self *Route) P(subpath string, route Route) {
	self.ParameterRoute(subpath, route)
}

func (self *Route) Route(subpath string, route Route) {
	self.route[subpath] = route
}

func (self *Route) R(subpath string, route Route) {
	self.Route(subpath, route)
}

type Router struct {
	route Route
}

func (self *Router) Parse(path string) (IController, []string){
	subpaths := strings.Split(path, "/")
	parameters := make([]string, 10)
	route := self.route
	for _, subpath := range subpaths {
		subroute, exist := route.route[subpath]
		if !exist {
			param := route.regex.Find(subpath)
			if len(param) > 0 {
				parameters = append(parameters, param)
				subroute = route.regexRoute
			}
		}
		if subroute == nil {
			break
		}
		route = subroute
	}
	controller := route.controller
	if controller == nil {
		RaiseHttp404(fmt.Errorf("Page not found"))
	}
	return controller, parameters
}