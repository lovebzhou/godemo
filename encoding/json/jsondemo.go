package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func bookDemo() {
	log.Println("================================================")
	type Book struct {
		Title       string   `json:"title"`
		Authors     []string `json:"autoers"`
		Publisher   string   `json:"publisher"`
		IsPublished bool     `json:"is_published"`
		Price       float32  `json:"price"`
	}

	book1 := Book{
		Title:       "Go语言编程",
		Authors:     []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
		Price:       9.99,
	}

	log.Printf("book1:%v\n", book1)

	b, err := json.Marshal(book1)

	book2 := Book{}

	if err := json.Unmarshal(b, &book2); err == nil {
		log.Printf("book2:%v", book2)
	}

	if err == nil {
		log.Printf("json:%v", string(b))
	}

	b3 := []byte(`{"title":"Go语言编程","autoers":["XuShiwei","HughLv","Pandaman","GuaguaSong","HanTuo","BertYuan","XuDaoli"],"publisher":"ituring.com.cn","is_published":true,"price":9.99}
	`)

	book3 := Book{}
	if err := json.Unmarshal(b3, &book3); err != nil {
		log.Printf("%s", err.Error())
	}
	log.Printf("book3:%v", book3)

	var data interface{}
	json.Unmarshal(b3, &data)
	log.Printf("data:%v", data)

	m := data.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			log.Println(k, "is bool", vv)
		case string:
			log.Println(k, "is string", vv)
		case float64:
			log.Println(k, "is float64", vv)
		case []interface{}:
			log.Println(k, "is an array:")
			for i, u := range vv {
				log.Println(i, u)
			}
		default:
			log.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func arbitraryData() {
	log.Println("================================================")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		log.Println(err)
		return
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			log.Println(k, "is string", vv)
		case float64:
			log.Println(k, "is float64", vv)
		case []interface{}:
			log.Println(k, "is an array:")
			for i, u := range vv {
				log.Println(i, u)
			}
		default:
			log.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func referenceTypes() {
	log.Println("================================================")

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}

	var m FamilyMember

	if err := json.Unmarshal(b, &m); err != nil {
		log.Println(err)
		return
	}

	log.Println(m)
}

// Address is
type Address struct {
	Type    string `json:"type"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// VCard is
type VCard struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Addresses []*Address `json:"addresses"`
	Remark    string     `json:"remark"`
}

func addressVCCard() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

func main() {
	bookDemo()
	arbitraryData()
	referenceTypes()
	addressVCCard()
}
