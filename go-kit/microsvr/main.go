package main

import (
	"flag"
	"log"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "http listen address")
		proxy =flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	var logger log.Logger
}