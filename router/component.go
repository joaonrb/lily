package router

import (
	"github.com/joaonrb/lily"
	"regexp"
)

const (
	charAmount  = 95
	charShift   = 32
	specialChar = `#`
)

//  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefgh
// ijklmnopqrstuvwxyz{|}~


// Node handles a step to the goal
type node struct {
	lily.Component
    nodes  [charAmount]lily.Component
}

func (n *node) Resolve(path []byte) interface{} {
	return n.nodes[path[0]-charShift].Resolve(path[1:])
}


type regexNode struct {
	node
	name   string
	next   *node
	regex  *regexp.Regexp
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

// end holds the treasure in the end of the route
type end struct {
	treasure interface{}
}

func (e *end) Resolve(path []byte) interface{} {
	return e.treasure
}


// Root is the first node for a route
type Root node

//TODO: Finish this
func (root *Root) Add(path []byte, treasure interface{}) error {
	current := root
	for i := 0; i < len(path); i++ {
        char := path[i]
        if char == specialChar {

		}
	}
}


func initNodes(nodes [charAmount]lily.Component) {
	for i := 0; i < charAmount; i++ {
		nodes[i] = EmptyComponentException
	}
}
