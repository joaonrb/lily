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

	"time"
)

const (
	DEFAULT_PORT        = 5555
	DEFAULT_HTTPS_PORT  = 443
	DEFAULT_BIND        = "127.0.0.1"
	DEFAULT_STATIC_PATH = "/static/"
)

func Run() {
	fmt.Printf("# Server Starting\n")

	if mainHandler == nil {
		mainHandler = defaultHandler()
	}

	port := Configuration.Port
	if port == 0 {
		if Configuration.Https {
			if Configuration.SSLCertificate == "" || Configuration.SSLKey == "" {
				fmt.Printf("**Error: Missing ssl certificate or key files in configuration.")
				os.Exit(1)
			}
			port = DEFAULT_HTTPS_PORT
		} else {
			port = DEFAULT_PORT
		}
	}
	bind := Configuration.Bind
	if bind == "" {
		bind = DEFAULT_BIND
	}

	read_timeout := time.Duration(Configuration.ReadTimeout * 10e6)
	write_timeout := time.Duration(Configuration.WriteTimeout * 10e6)

	for _, middleware := range Configuration.Middleware {
		resgistedMiddleware[middleware](mainHandler)
	}

	mux := http.NewServeMux()
	if Configuration.StaticFiles != "" {
		if Configuration.StaticPath == "" {
			mux.Handle(DEFAULT_STATIC_PATH, http.FileServer(http.Dir(Configuration.StaticFiles)))
		} else {
			mux.Handle(Configuration.StaticFiles, http.FileServer(http.Dir(Configuration.StaticFiles)))
		}
	}
	if Configuration.StaticPath != "/" {
		mux.Handle("/", mainHandler)
	}


	address := fmt.Sprintf("%s:%d", bind, port)
	server := &http.Server{
		Addr: address,
		Handler: mux,
		ReadTimeout: read_timeout,
		WriteTimeout: write_timeout,
	}
	go func() {
		var err error
		if Configuration.Https {
			server.ListenAndServeTLS(Configuration.SSLCertificate, Configuration.SSLKey)
		} else {
			server.ListenAndServe()
		}
		if err != nil {
			fmt.Printf("**Error starting server**\n%s\n\nExiting. Bye bye...", err.Error())
			os.Exit(1)
		}
	}()
	fmt.Printf("# Listening at %s\n", address)
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