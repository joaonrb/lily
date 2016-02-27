//
// Copyright (c) João Nuno. All rights reserved.
//
package main

import (
	"net/http"
	"lily"
	"lily/apps/accesslog"
	"lily/examples/hello"
)

func main() {
	lily.Configuration = &lily.Settings{
		Loggers: []lily.LogSettings{
			{
				Type: "console",
				Layout: "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}",
				Level: "debug",
			},
		},
		AccessLog: lily.AccessLogSetings{
			Type: "console",
		}, 
	}
	lily.LoadLogger()
	
	controller := &hello.HelloWorldController{}

	lily.Register([]lily.Way{
		{"/", controller},
	})
	
	handler := lily.NewHandler(
		lily.NewRequestInitializer(),
		lily.NewFinalizer(),
	)
	
	accesslog.Register(handler)
	
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}