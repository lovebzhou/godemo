package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello:1")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello:2")

	fmt.Println("Hello go!")
}
