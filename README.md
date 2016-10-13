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
To install Lily, make sure you have installed [Go 1.5](https://storage.googleapis.com/golang/go1.5.src.tar.gz) or [later
version](https://storage.googleapis.com/golang/go1.7.src.tar.gz) correctly.

First get the dependencies:

- [go-logging](https://github.com/op/go-logging)
- ~~[yalm](https://gopkg.in/yaml.v2)~~
- ~~[macaron](https://github.com/go-macaron/cache)   # For cache~~
- [fasthttp](https://github.com/valyala/fasthttp)  # Fast and simple web tools

Then get lily.

```
$ go get github.com/op/go-logging
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
	lily.BaseController
}

func (self *HelloWorldController) Get(request *fasthttp.RequestCtx, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = "<h1>Hello World!</h1>"
	return response
}

// Controller that receive the parameters "user" and say hello to it.
type RegexHelloWorldController struct {
	lily.BaseController
}
func (self *RegexHelloWorldController) Get(request *fasthttp.RequestCtx, args map[string]string) *lily.Response {
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

	lily.Url("/", controller)
	lily.Url("/another", controller)
	lily.Url(`/:(?P<user>\S+)`, regexController)
}
```

The router will parse the paths and build a tree with the flat names and a list with the regular expressions. The
regex parts **must** start with ":". This will make the parser interpret the rest as a regex.

The last part we need the main. The main file must exist outside the package of the app.

```
cd ..
```

```go
//hello-server.go
package main

import (
	"github.com/joaonrb/lily"
	_ "hello"
	"fmt"
	"os"
)

func main() {
    h := fasthttp.CompressHandler(lily.CoreHandler)
	if err := fasthttp.ListenAndServe(":8080", h); err != nil {
    	log.Fatalf("Error in ListenAndServe: %s", err)
    }
}
```

The main is very simple. It have to get the config file. Is very important that the middleware that is being used is
imported somewhere in the code. Since we did not import access log and we need it to initialize we have to import it
here. Also we have to import the hello package and then lily take care the rest.

```
$ go run hello-server.go
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


