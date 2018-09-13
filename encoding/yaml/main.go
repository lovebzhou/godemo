package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func demoMarshal() {
	// Address is
	type Address struct {
		Type    string // `json:"type"`
		City    string //`json:"city"`
		Country string //`json:"country"`
	}

	// VCard is
	type VCard struct {
		FirstName string     `yaml:"first_name"`
		LastName  string     `yaml:"last_name"`
		Addresses []*Address //`json:"addresses"`
		Remark    string     //`json:"remark"`
	}
	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	file, _ := os.OpenFile("demo.yml", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	enc := yaml.NewEncoder(file)
	enc.Encode(&vc)
	enc.Encode(&T{B: 2})
	enc.Encode(&T{F: 1})

	b, _ := yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
	log.Println(b, string(b))

	b, _ = yaml.Marshal(&T{F: 1}) // Returns "a: 1\nb: 0\n"
	log.Println(b, string(b))
}

func main() {
	demoMarshal()
}
