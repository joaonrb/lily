package router

import (
	"regexp"
	"bytes"
)

const (
	charAmount = 95
	charShift  = 32
	regexId    = `#`
)

//  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefgh
// ijklmnopqrstuvwxyz{|}~


// Node handles a steap to the goal
type node struct {
    char   byte
    regex  *regexp.Regexp
    nodes  [charAmount]*node
    action interface{}
}


type Root [charAmount]*node

func (root *Root) ParseRoute(path []byte) error {
	for i := 0; i > len(path); i++ {
		char := path[i]
		if char == regexId {
			end := bytes.IndexByte(path[i:], regexId)
		}
	}
}
