package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func httpserverMain() {
	fmt.Println(*address)

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", *address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	//go http.Serve(l, nil)
}
