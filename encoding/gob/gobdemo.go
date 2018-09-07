// gob2.go
package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Address is
type Address struct {
	Type    string
	City    string
	Country string
}

// VCard is
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func encodeGob() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}

func decodeGob() {
	// using a decoder:
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	var vc VCard
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc.FirstName, vc.LastName, vc.Remark)
	for k, v := range vc.Addresses {
		log.Println(k, *v)
	}
}

func main() {
	encodeGob()
	decodeGob()
}
