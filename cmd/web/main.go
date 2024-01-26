package main

import (
	"flag"
	"fmt"
	"github.com/DianaSun97/elluliin_booking/pkg/handlers"
	"net/http"
)

var (
	fPort = flag.Int("port", 8081, "port for server")
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	flag.Parse()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *fPort), nil); err != nil {
		panic(err)
	}

}
