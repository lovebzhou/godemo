package main

import (
	"flag"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net"
	"net/http"
)

var host *string = flag.String("h", "localhost", "host")
var port *string = flag.String("p", "8080", "port")

func main() {
	flag.Parse()
	addr := *host + ":" + *port
	log.Printf("listen on %s", addr)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/lookup/#host", func(w rest.ResponseWriter, req *rest.Request) {
			ip, err := net.LookupIP(req.PathParam("host"))
			if err != nil {
				rest.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteJson(&ip)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(addr, api.MakeHandler()))
}
