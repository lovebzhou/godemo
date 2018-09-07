package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	b := []int{1, 2, 3, 4}
	fmt.Println(b)

	fmt.Println("b.length=", len(b), "b.capacity=", cap(b))
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)
}
