package main

import (
	"fmt"
	"log"
)

// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer log.Printf("defer: i = %d", i)
	i++
	log.Printf("testDefer1: i = %d", i)
}

// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	for i := 0; i < 4; i++ {
		defer log.Printf("defer: i = %d", i)
		log.Printf("i = %d", i)
	}
}

// 3. Deferred functions may read and assign to the returning function's named return values.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()
	log.Println("Calling g.")
	g(0)
	log.Println("Returned normally from g.")
}

func g(i int) {
	defer func() {
		if i == 3 {
			if r := recover(); r != nil {
				log.Println("Recovered in g", r)
				panic(fmt.Sprintf("r = %v, r = %v", i, r))
			}
		}
	}()

	if i > 3 {
		log.Println("Panicking!")
		// panic(fmt.Sprintf("%v", i))
		panic(i)
	}
	defer log.Println("Defer in g", i)
	log.Println("Printing in g", i)
	g(i + 1)
}

func main() {
	fmt.Printf("=============================== defer ==============================\n\n")
	a()
	b()
	i := c()
	log.Printf("i = c() = %d", i)

	fmt.Printf("\n=================== defer + panic + recover =======================\n\n")
	f()
	log.Println("Returned normally from f.")
}
