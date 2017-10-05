package router

import (
	"bytes"
	"github.com/joaonrb/lily"
	"regexp"
	"fmt"
)

const (
	charAmount        = 78  // 119 // 95
	charShift         = 45  // 10  // 32
	specialChar       = 64  // Character @ 36  // Character #
	scapeChar         = 63  // Character ? 10  // Character \n
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
	add(append(path, scapeChar), n, treasure)
}

func (n *node) String() string {
	return fmt.Sprintf("Route node '%s'", string(n.char))
}

type regexContainer struct {
	regex     *regexp.Regexp
	component lily.Component
}

type regexNode struct {
	node
	next  lily.Component
	regex []*regexContainer
}

func (rn *regexNode) Resolve(context interface{}) interface{} {
	path := context.([]byte)
	// TODO: recycle this match with sync.Pool
	for _, regex := range rn.regex {
		match := regex.regex.Find(path)
		if len(match) != 0 {
			// TODO: add parameters to the result. Wrap them on struct
			return regex.component.Resolve(path[len(match)-charShift:])
		}
	}
	return rn.node.Resolve(path)
}

func (rn *regexNode) Add(path []byte, treasure lily.Component) {
	add(append(path, scapeChar), rn, treasure)
}

func (rn *regexNode) String() string {
	return fmt.Sprintf("Route regex node '%s'", string(rn.char))
}

// end holds the component in the end of the route
type end struct {
	lily.Component
	char      byte
	component lily.Component
}

func (e *end) Resolve(interface{}) interface{} {
	return e.component
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

func add(path []byte, self, c lily.Component) lily.Component {
	if path[0] < charShift || path[0] > charShift+charAmount {
		panic(ThrowInvalidCharacterException(
			fmt.Sprintf("Character %s is not allowed", string(path[0])),
		))
	}
	newNode := getNode(path, self, c)
	switch self := self.(type) {
	case *node:
		self.nodes[path[0]-charShift] = newNode
	case *regexNode:
		self.next = newNode
	}
	return self
}

func getNode(path []byte, self, c lily.Component) lily.Component {
    char, rest := path[0], path[1:]
	switch char {
	case scapeChar:
		return &end{char: char, component: c}
	case specialChar:
		return initRegex(rest, self, c)
	default:
		var n lily.Component
		switch self := self.(type) {
		case *node:
			n = self.nodes[char-charShift]
		}
		if n != EmptyComponentException {
			return add(rest, n, c)
		}
		return add(rest, &node{char: char, nodes: initNodes()}, c)
	}
}

func initRegex(path []byte, self, c lily.Component, ) lily.Component {
	i := bytes.IndexByte(path[1:], regexParserFormat)
	regex, rest := path[1:i-1], path[i+1:]
	newNode := getNode(rest, self, c)
	switch self := self.(type) {
	case *node:
		return &regexNode{
			node:  *self,
			regex: []*regexContainer{
				&regexContainer{
					regexp.MustCompile(string(regex)),
					newNode,
				},
			},
		}
	case *regexNode:
		 self.regex = append(self.regex, &regexContainer{
			regexp.MustCompile(string(regex)),
			newNode,
		})
		return self
	}
	return nil
}
