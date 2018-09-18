package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func demo1() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

// get from https server is same as http one
func demo2() {
	resp, err := http.Get("https://lovebzhou.net:8443/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

// For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
func demo3() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

func main() {
	demo1()
	demo2()
	demo3()
}
