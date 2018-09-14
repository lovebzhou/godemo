package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Arguments struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Arguments, reply *int) error {
	*reply = args.A * args.B
	log.Printf("A=%d, B=%d, R=%d", args.A, args.B, *reply)
	return nil
}

func (t *Arith) Divide(args *Arguments, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divive by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	log.Printf("A=%d, B=%d, Quo=%d, Rem=%d\n", args.A, args.B, quo.Quo, quo.Rem)
	return nil
}

var host *string = flag.String("h", "localhost", "host")
var port *string = flag.String("p", "8088", "port")

func main() {
	flag.Parse()

	//address := fmt.Sprintf("%v:%v", *host, *port)
	address := *host + ":" + *port
	fmt.Println(address)

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", address)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	//go http.Serve(l, nil)
}
