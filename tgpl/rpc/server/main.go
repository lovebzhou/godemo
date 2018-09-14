package main

import (
	"flag"
	"fmt"
	"os"
)

var transform = flag.String("t", "http", "transform type")
var address = flag.String("address", ":8088", "host:port")

func main() {
	flag.Parse()

	switch *transform {
	case "http":
		httpserverMain()
	case "tcp":
		tcpserverMain()
	case "json":
		jsonserverMain()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
