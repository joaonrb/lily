//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package hello

import (
	"github.com/joaonrb/lily"
)

func init() {
	controller := &HelloWorldController{}
	regexController := &RegexHelloWorldController{}

	lily.RegisterRoute([]lily.Way{
		{"/", controller},
		{"/another", controller},
		{`/:(?P<user>\S+)`, regexController},
	})
}
