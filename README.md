Lily
====
[![Build Status](https://travis-ci.org/joaonrb/lily.svg?branch=master)](https://travis-ci.org/joaonrb/lily)
[![codecov.io](https://codecov.io/github/joaonrb/lily/coverage.svg?branch=master)](https://codecov.io/github/joaonrb/lily?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/joaonrb/lily)](https://goreportcard.com/report/github.com/joaonrb/lily)

Plan
----
Join to Lily plan group in [Trello Group](https://trello.com/invite/lily745/4a8efdb4ab693b2aa263546c8380c249).

About
-----
The **Lily** is a webframework(still very simple) using Go language. In this stage is more a wrapper of very good tools
for web development organized in VMC paradigm.

Lily was inspired in Django framework. In particular the way contexts may flow between diferent stages of the request
processing before arrive and after leaving the controller. Throught middlewares, Django make this behaviour very
pluggable helping the Django Apps fairly independent to include in multiple projects.

I also give it a try to a simple, funcional and, most important, readable routing system.

Installation
------------
To install Lily, make sure you have installed [Go 1.5](https://storage.googleapis.com/golang/go1.6.src.tar.gz) or later
version correctly.

First get the dependencies:

- [go-logging](https://github.com/op/go-logging)
- [yalm](https://gopkg.in/yaml.v2)
- [macaron](https://github.com/go-macaron/cache)   # For cache
- [fasthttp](https://github.com/valyala/fasthttp)  # Fast and simple web tools

Then get lily.

```
$ go get github.com/op/go-logging
$ go get gopkg.in/yaml.v2
$ go get github.com/go-macaron/cache  # For Caching
$ go get github.com/valyala/fasthttp  # For the fast stuff
$ go get github.com/joaonrb/lily
```

Getting Started
---------------
Like Django, to start serving content with Lily you will need a Lily app. The app has controllers and middlewares. Is
not mandatory for the app to work to have this two components, but if you want it to do anything is better to have at
least one of this.

We ca start by creating the app folder.
```
$ mkdir hello
$ cd hello
```

Then we can create some controllers.
```go
// hello/controllers.go
package hello

import (
	"github.com/joaonrb/lily"
	"fmt"
)

// Simple controller that say hello to the worlds
type HelloWorldController struct {
	lily.Controller
}

func (self *HelloWorldController) Get(request *lily.Request, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = "<h1>Hello World!</h1>"
	return response
}

// Controller that receive the parameters "user" and say hello to it.
type RegexHelloWorldController struct {
	lily.Controller
}
func (self *RegexHelloWorldController) Get(request *lily.Request, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = fmt.Sprintf("<h1>Hello %s!</h1>", args["user"])
	return response
}
```

We can create also a router file. The routing structure doesn't have to live in an independent file. It can libe in
the main or any other go file with your code. Your choice. I prefer to have a file just for it.

```go
// hello/router.go
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
```

The router will parse the paths and build a tree with the flat names and a list with the regular expressions. The
regex parts **must** start with ":". This will make the parser interpret the rest as a regex.

The last part we need the main. The main file must exist outside the package of the app.

```
cd ..
```

First we need to make a config file. The configuration file is done in yaml. The options are still very litle and thers
just so many to do.

```yaml
# hello-config.yaml

# Lily hello world configuration file

# The address for the server to bind.
# Default: 127.0.0.1
bind: 0.0.0.0

# Port for the server to listen for.
# Default: 5555
port: 8000

# Here you can register the middleware you want to use in your app.
# To the name here to be recognised by the middleware installer it must be registered
# and for that it have to be imported somewhere in your project or manually resisted
# using the lily.RegisterMiddleware function.
middlewares:
  - accesslog

# Loggers
# More than one logger can be deployed. The names of the loggers can be any string as long
# as they don't repeat.
# Lily log implementation is on top of go-logging. Details on the layout format can be found
# here - http://github.com/op/go-logging.
loggers:
  main:
    type:   console
    layout: "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
    level:  debug
  file:
    type:   file
    path:   /tmp/linda.log
    layout: "%{time:2006-01-02 15:04:05.000} %{level:.4s} %{shortfile} %{message}"
    level:  info

# The access log is only one and can be showed in the console or in file. In case of
# using file an extra parameter for path is in order.
accesslog:
  type: console

```

```go
//hello-server.go
package main

import (
	"github.com/joaonrb/lily"
	_ "github.com/joaonrb/lily/apps/accesslog"
	_ "hello"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <config>\n", os.Args[0])
		os.Exit(1)
	}
	// Pass the absolute path of the file hello-config.yaml in command
	err := lily.Init(os.Args[1])
	if err != nil {
		fmt.Printf("Errors in config file: %s\n", err.Error())
	}

	lily.LoadLogger()
	lily.Run()
}
```

The main is very simple. It have to get the config file. Is very important that the middleware that is being used is
imported somewhere in the code. Since we did not import access log and we need it to initialize we have to import it
here. Also we have to import the hello package and then lily take care the rest.

```
$ go run hello-server.go hello-config.yaml
# Server Starting
# Listening at 0.0.0.0:8000
# Use Ctrl+C to close
```

That's it. You have a fully working go server that say hello to people. You can try in your browser

- http://localhost:8000
- http://localhost:8000/another
- http://localhost:8000/girl

Coverage
--------

![codecov.io](https://codecov.io/github/joaonrb/lily/branch.svg?branch=master)

Wishlist
--------

- Fully covered with unitests
- Nice forms Django style


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/joaonrb/lily/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

