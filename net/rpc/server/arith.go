package main

import (
	"errors"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	log.Printf("A=%d, B=%d, R=%d", args.A, args.B, *reply)
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divive by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	log.Printf("A=%d, B=%d, Quo=%d, Rem=%d\n", args.A, args.B, quo.Quo, quo.Rem)
	return nil
}
