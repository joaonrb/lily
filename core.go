//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"os"
	"os/signal"
	"fmt"
	"time"
	"github.com/valyala/fasthttp"
)

const (
	DEFAULT_PORT        = 5555
	DEFAULT_HTTPS_PORT  = 443
	DEFAULT_BIND        = "127.0.0.1"
)

func Run() {
	fmt.Printf("# Server Starting\n")

	if mainHandler == nil {
		mainHandler = defaultHandler()
	}

	port := Configuration.Port
	if port == 0 { port = DEFAULT_PORT }

	var httpsPort int
	if Configuration.Https {
		if Configuration.SSLCertificate == "" || Configuration.SSLKey == "" {
			fmt.Printf("**Error: Missing ssl certificate or key files in configuration.")
			os.Exit(1)
		}
		httpsPort = Configuration.HttpsPort
		if httpsPort == 0 { httpsPort = DEFAULT_HTTPS_PORT }
	}
	bind := Configuration.Bind
	if bind == "" { bind = DEFAULT_BIND }

	read_timeout := time.Duration(Configuration.ReadTimeout * 10e6)
	write_timeout := time.Duration(Configuration.WriteTimeout * 10e6)

	// Register middleware
	for _, middleware := range Configuration.Middleware {
		resgistedMiddleware[middleware](mainHandler)
	}

	for uri, path := range Configuration.StaticFiles {
		mainHandler.RegisterStaticPath(uri, path)
	}

	address := fmt.Sprintf("%s:%d", bind, port)

	server := &fasthttp.Server{
		Handler: mainHandler.ServeHTTP,
		Name: "Lily Server",
		ReadTimeout: read_timeout,
		WriteTimeout: write_timeout,
	}
	go func() {
		var err error
		switch {
		case Configuration.Https:
			err = server.ListenAndServeTLS(fmt.Sprintf("%s:%d", bind, httpsPort), Configuration.SSLCertificate,
				Configuration.SSLKey)
		case Configuration.UnixSocket:
			err = server.ListenAndServeUNIX(bind, os.ModePerm)
		default:
			err = server.ListenAndServe(fmt.Sprintf("%s:%d", bind, port))
		}
		if err != nil {
			fmt.Printf("\n**Error starting server**\n<<%s>>\n\nExiting. Bye bye...\n", err.Error())
			os.Exit(1)
		}
	}()
	fmt.Printf("# Listening at %s\n", address)
	fmt.Printf("# Use Ctrl+C to close \n")

	waitForFinish()
}

// Code rented from http://nathanleclaire.com/blog/2014/08/24/handling-ctrl-c-interrupt-signal-in-golang-programs/
func waitForFinish() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\n# Server closing...\n")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}