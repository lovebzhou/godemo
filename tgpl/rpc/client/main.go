package main

import "flag"

var transform = flag.String("t", "http", "transform type")
var address = flag.String("address", ":8088", "host:port")

func main() {
	flag.Parse()

	switch *transform {
	case "http":
		httpclientMain()
	case "tcp":
		tcpclientMain()
	case "json":
		jsonclientMain()
	}
}
