package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	fPort = flag.Int("port", 8080, "port for server")
)

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%d", *fPort), nil)
}
