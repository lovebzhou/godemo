package main

import (
	"flag"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

var host *string = flag.String("h", "localhost", "host")
var port *string = flag.String("p", "8080", "port")

func main() {
	log.Printf("Listen at %s:%s", *host, *port)
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(*host+":"+*port, api.MakeHandler()))
}
