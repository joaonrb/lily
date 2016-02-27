//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package main

import (
	"lily"
	"lily/examples/hello"
	_ "lily/apps/accesslog"
)

func main() {
	lily.Configuration = &lily.Settings{
		Middleware: []string{"accesslog"},
		Loggers: map[string]lily.LogSettings{
			"default": {
				Type: "console",
				Layout: "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}",
				Level: "debug",
			},
		},
		AccessLog: lily.AccessLogSettings{
			Type: "console",
		}, 
	}
	lily.LoadLogger()
	
	controller := &hello.HelloWorldController{}

	lily.RegisterRoute([]lily.Way{
		{"/", controller},
		{"/test", controller},
		{"/test/r", controller},
	})

	lily.Run()
}