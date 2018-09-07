package main

import "fmt"

func fab(n int) {
	a, b := 0, 1
	for a < n {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func main() {
	fab(200)
}
