package router

import (
	"bytes"
	"github.com/joaonrb/lily"
	"regexp"
	"fmt"
)

const (
	charAmount        = 119 // 95
	charShift         = 10  // 32
	specialChar       = 36  // Character #
	scapeChar         = 10  // Character \n
	regexParserFormat = 96  // Character `
)

// Node handles a step to the goal
type node struct {
	lily.Component
	char  byte
	nodes [charAmount]lily.Component
}

func (n *node) Resolve(context interface{}) interface{} {
	path := context.([]byte)
	return n.nodes[path[0]-charShift].Resolve(path[1:])
}

func (n *node) Add(path []byte, treasure lily.Component) {
	add(n, append(path, scapeChar), treasure)
}

func (n *node) String() string {
	return fmt.Sprintf("Route node '%s'", string(n.char))
}

type regexNode struct {
	node
	name  string
	next  lily.Component
	regex *regexp.Regexp
}

func (rn *regexNode) Resolve(context interface{}) interface{} {
	path := context.([]byte)
	// TODO: recycle this match with sync.Pool
	match := rn.regex.Find(path)
	if len(match) != 0 {
		// TODO: add parameters to the result. Wrap them on struct
		return rn.next.Resolve(path[len(match)-charShift:])
	}
	return rn.node.Resolve(path)
}

func (rn *regexNode) Add(path []byte, treasure lily.Component) {
	add(rn, append(path, scapeChar), treasure)
}

func (rn *regexNode) String() string {
	return fmt.Sprintf("Route regex node '%s'", string(rn.char))
}

// end holds the treasure in the end of the route
type end struct {
	lily.Component
	char     byte
	treasure lily.Component
}

func (e *end) Resolve(interface{}) interface{} {
	return e.treasure
}

func (e *end) String() string {
	return fmt.Sprintf("Route end '%s'", string(e.char))
}

// Root is the first node for a route
type Root struct{
	node
}

func (r *Root) Resolve(context interface{}) interface{} {
	path := context.([]byte)
	return r.node.Resolve(append(path, scapeChar))
}

func (*Root) String() string {
	return fmt.Sprintf("Route root")
}

func New() *Root {
	return &Root{node{char: charShift, nodes: initNodes()}}
}

func initNodes() [charAmount]lily.Component {
	nodes := [charAmount]lily.Component{}
	for i := 0; i < charAmount; i++ {
		nodes[i] = EmptyComponentException
	}
	return nodes
}

func add(self lily.Component, path []byte, treasure lily.Component) {
	newNode, rest := getNode(self, path, treasure)
	if len(rest) != 0 {
		add(newNode, rest, treasure)
	}
	switch self := self.(type) {
	case *node:
		self.nodes[path[0]-charShift] = newNode
	case *regexNode:
		self.next = newNode
	}
}

func getNode(self lily.Component,
	path []byte, treasure lily.Component) (lily.Component, []byte) {

	switch path[0] {
	case scapeChar:
		return &end{char: path[0], treasure: treasure}, nil
	case specialChar:
		return initRegex(path[1:])
	default:
		return &node{char: path[0], nodes: initNodes()}, path[1:]
	}
}

//func checkNode(self lily.Component, char byte) lily.Component {
//
//
//	switch self := self.(type) {
//	case *node:
//		return self.nodes[char]
//	case *regexNode:
//		self.next = newNode
//	}
//}

func initRegex(path []byte) (lily.Component, []byte) {
	i := bytes.IndexByte(path[1:], regexParserFormat)
	regex, rest := path[1:i-1], path[i+1:]
	newNode := &regexNode{
		node:  node{char:charShift, nodes: initNodes()},
		regex: regexp.MustCompile(string(regex)),
	}
	if len(newNode.regex.SubexpNames()) > 1 {
		newNode.name = newNode.regex.SubexpNames()[1]
	}
	return newNode, rest
}
