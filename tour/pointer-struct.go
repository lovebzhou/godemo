package main

import "fmt"
import "strings"

type Vertex struct {
	X int
	Y int
	z int
}

func main() {
	v := Vertex{1, 2, 3}
	p := &v
	p.X = 1e9
	p.z = 30
	fmt.Println(v)
	fmt.Println(len("dd"))
}
