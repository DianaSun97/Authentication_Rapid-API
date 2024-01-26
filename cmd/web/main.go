package main

import (
	"flag"
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

var (
	fPort = flag.Int("port", 8080, "port for server")
)

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%d", *fPort), nil)
}
