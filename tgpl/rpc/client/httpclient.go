package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func httpclientMain() {
	client, err := rpc.DialHTTP("tcp", *address)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replyCall := <-divCall.Done

	fmt.Println(replyCall)
	fmt.Printf("quo:%d, rem:%d\n", quotient.Quo, quotient.Rem)
}
