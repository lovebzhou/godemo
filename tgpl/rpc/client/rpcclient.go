package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Arguments struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var host *string = flag.String("h", "localhost", "host")
var port *string = flag.String("p", "8088", "port")

func main() {
	/*
		if len(os.Args) != 2 {
			log.Fatal("Usage: %s host:port", os.Arguments[0])
		}
	*/

	fmt.Printf("Lenght of Args = %d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Println(i, v)
	}

	flag.Parse()
	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Printf("address=%s\n", address)

	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Arguments{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done

	fmt.Println(replyCall)
	fmt.Printf("quo:%d, rem:%d\n", quotient.Quo, quotient.Rem)
}
