package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type myHandler struct {
	context string
}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<p>context: %s</p><p>path: %s</p>", h.context, r.URL.Path)
}

func mainListenAndServe() {
	http.Handle("/foo", myHandler{context: "/foo"})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h2>Hello, %q</h2>", html.EscapeString(r.URL.Path))
	})

	// log.Fatal(http.ListenAndServe(":8080", myHandler{context: "global handler"}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
