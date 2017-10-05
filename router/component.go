package router

import (
	"bytes"
	"github.com/joaonrb/lily"
	"regexp"
)

const (
	charAmount        = 95
	charShift         = 32
	specialChar       = 36 // Character #
	scapeChar         = 10 // Character \n
	regexParserFormat = 96 // Character `
)

// Node handles a step to the goal
type node struct {
	lily.Component
	nodes [charAmount]lily.Component
}

func (n *node) Resolve(path []byte) interface{} {
	return n.nodes[path[0]-charShift].Resolve(path[1:])
}

func (n *node) Add(path []byte, treasure lily.Component) {
	add(n, append(path, scapeChar), treasure)
}

type regexNode struct {
	node
	name  string
	next  lily.Component
	regex *regexp.Regexp
}

func (rn *regexNode) Resolve(path []byte) interface{} {
	// TODO: recycle this match with sync.Pool
	match := rn.regex.Find(path)
	if len(match) != 0 {
		// TODO: add parameters to the result. Wrap them on struct
		return rn.next.Resolve(path[len(match):])
	}
	return rn.node.Resolve(path)
}

func (rn *regexNode) Add(path []byte, treasure lily.Component) {
	add(rn, append(path, scapeChar), treasure)
}

// end holds the treasure in the end of the route
type end struct {
	lily.Component
	treasure lily.Component
}

func (e *end) Resolve(path []byte) lily.Component {
	return e.treasure
}

// Root is the first node for a route
type Root node

func initNodes() [charAmount]lily.Component {
	nodes := [charAmount]lily.Component{}
	for i := 0; i < charAmount; i++ {
		nodes[i] = EmptyComponentException
	}
	return nodes
}

func add(self lily.Component, path []byte, treasure lily.Component) {
	newNode, rest := getNode(path, treasure)
	add(newNode, rest, treasure)
	switch self := self.(type) {
	case *node:
		self.nodes[path[0]] = newNode
	case *regexNode:
		self.next = newNode
	}
}

func getNode(path []byte, treasure lily.Component) (lily.Component, []byte) {
	switch path[0] {
	case scapeChar:
		return &end{treasure: treasure}, path[1:]
	case specialChar:
		return initRegex(path[1:])
	default:
		return node{nodes: initNodes()}, path[1:]
	}
}

func initRegex(path []byte) (lily.Component, []byte) {
	i := bytes.IndexByte(path[1:], regexParserFormat)
	regex, rest := path[1:i-1], path[i+1:]
	newNode := &regexNode{
		node:  node{nodes: initNodes()},
		regex: regexp.MustCompile(string(regex)),
	}
	if len(newNode.regex.SubexpNames()) > 1 {
		newNode.name = newNode.regex.SubexpNames()[1]
	}
	return newNode, rest
}
