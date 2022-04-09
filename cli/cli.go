package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/ssurepa/go_blockchain/explorer"
	"github.com/ssurepa/go_blockchain/rest"
)

func usage() {
	fmt.Printf("Welcome to JO Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: Set the PORT of the server\n")
	fmt.Printf("-mode: Choose 'html, 'rest', or 'both\n'both' will run second server on port+1000")
	runtime.Goexit()
}

func Start() {

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1000)
	default:
		usage()
	}
}
