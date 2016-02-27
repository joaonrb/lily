//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"net/http"
	"os"
	"os/signal"
	"fmt"

)

const (
	DEFAULT_PORT  = 5555
	DEFAULT_BIND  = "127.0.0.1"
)

func Run() {
	fmt.Printf("# Server Starting\n")

	if mainHandler == nil {
		mainHandler = defaultHandler()
	}

	port := Configuration.Port
	if port == 0 {
		port = DEFAULT_PORT
	}
	bind := Configuration.Bind
	if bind == "" {
		bind = DEFAULT_BIND
	}

	for _, middleware := range Configuration.Middleware {
		resgistedMiddleware[middleware](mainHandler)
	}

	http.Handle("/", mainHandler)
	listener := fmt.Sprintf("%s:%d", bind, port)
	go http.ListenAndServe(listener, nil)
	fmt.Printf("# Listening at %s\n", listener)
	fmt.Printf("# Use Ctrl+C to close\n")

	waitForFinish()
}


// Code rented from http://nathanleclaire.com/blog/2014/08/24/handling-ctrl-c-interrupt-signal-in-golang-programs/
func waitForFinish() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\n# Server closing...")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}