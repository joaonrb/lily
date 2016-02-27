//
// Author João Nuno.
//
// Default router for Lily. Load the router in main app. Router must implement IRouter.
// Package router loads a string in format:
// "/path1/path2/:(?P<parameter>\d+)" ExampleController
//
package lily

import (
	"regexp"
	"strings"
	"fmt"
)

var mainRouter IRouter

// Register a new router
func ForceRegisterRouter(router IRouter) {
	mainRouter = router
}

// Try to register a new router if no router is register already.
// @return: true if new router was successful registered and false otherwise.
func RegisterRouter(router IRouter) bool {
	if mainRouter == nil {
		ForceRegisterRouter(router)
		return true
	}
	return false
}

// Register a path to the router. If no router was resisted it creates a new router.
func RegisterPath(path string, controller IController) error {
	// Creates a new router if none exist.
	RegisterRouter(&Router{newRouterNode()})
	return mainRouter.Register(path, controller)
}

// Register a bulk of controllers
func RegisterRoute(paths []Way) error {
	// Creates a new router if none exist.
	RegisterRouter(&Router{newRouterNode()})

	for _, way := range paths {
		err := mainRouter.Register(way.Path, way.Controller)
		if err != nil {
			return err
		}
	}
	return nil
}

type IRouter interface {
	Parse(string) (IController, map[string]string, error)
	Register(string, IController) error
}

// Router node
// It holds the possible flat routes and after that the regex routes.
type routerNode struct {
	flatRoutes   map[string]*routerNode
	regexRoutes  map[string]*regexNode
	controller   IController
}

func newRouterNode() *routerNode {
	return &routerNode{map[string]*routerNode{}, map[string]*regexNode{}, nil}
}

type regexNode struct {
	*routerNode
	regex  *regexp.Regexp
}
// Simple router
// Search for the flat routes first and the regex after. At the end returns the controller.
type Router struct {
	route *routerNode
}

type Way struct {
	Path        string
	Controller  IController
}

func (self *Router) Parse(path string) (IController, map[string]string, error) {
	ways := strings.Split(path, "/")
	thisRoute := self.route
	params := map[string]string{}
	fmt.Println(self.route.controller == nil, path, ways[1:len(ways)-1])
	for _, way := range ways[1:len(ways)-1] {
		if newRoute, ok := thisRoute.flatRoutes[way]; ok {
			thisRoute = newRoute
		} else {
			found := false
			for _, regexRoute := range thisRoute.regexRoutes {
				match := regexRoute.regex.FindStringSubmatch(way)
				if len(match) > 0 {
					params[regexRoute.regex.SubexpNames()[0]] = match[0]
					thisRoute = regexRoute.routerNode
					found = true
					break
				}
			}
			if !found {
				return nil, nil, NewHttp404()
			}
		}
	}
	fmt.Println(thisRoute.controller, params, nil)
	return thisRoute.controller, params, nil
}

func (self *Router) Register(path string, controller IController) error {
	ways := strings.Split(path, "/")
	thisRoute := self.route

	for _, way := range ways {
		switch {
		case len(way) == 0:
		case ':' == way[0]:
			regexString := way[1:]
			if regexRouter, ok := thisRoute.regexRoutes[regexString]; ok {
				thisRoute = regexRouter.routerNode
			} else {
				regex, err := regexp.Compile(regexString)
				if err != nil {
					return err
				}
				thisRoute.regexRoutes[regexString] = &regexNode{newRouterNode(), regex}
				thisRoute = thisRoute.regexRoutes[regexString].routerNode
			}
		default:
			if router, ok :=  thisRoute.flatRoutes[way]; ok {
				thisRoute = router
			} else {
				thisRoute = newRouterNode()
			}
		}
	}
	switch {
	case thisRoute.controller != nil:
		return NewPathAlreadyExist(path)
	default:
		thisRoute.controller = controller
	}
	return nil
}