//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package main

import (
	"lily"
	_ "lily/examples/hello"
	_ "lily/apps/accesslog"
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