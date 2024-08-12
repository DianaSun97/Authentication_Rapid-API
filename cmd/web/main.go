package main

import (
	"flag"
	"fmt"
	"github.com/DianaSun97/elluliin_booking/pkg/config"
	"github.com/DianaSun97/elluliin_booking/render"
	"log"
	"net/http"
)

var (
	fPort = flag.Int("port", 8081, "port for server")
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	app.UseCache = false
	if err != nil {
		log.Fatalln("cannot create template cach")
	}

	app.TemplateCache = tc
	//http.HandleFunc("/", handlers.Home)
	//http.HandleFunc("/about", handlers.About)

	flag.Parse()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *fPort), nil); err != nil {
		panic(err)
	}

}
