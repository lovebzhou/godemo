package main

import "fmt"

func _fab1() func() int {
	a, b := 0, 1
	return func() int {
		r := a
		a, b = b, a+b
		return r
	}
}

func fab1(n int) {
	f := _fab1()
	for r := f(); r < n; r = f() {
		fmt.Println(r)
	}
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func acc() func(int) int {
	x := 0
	return func(delta int) int {
		x += delta
		return x
	}
}

func main() {
	fab1(120)

	add2 := add2()
	add3 := adder(3)
	for i := 0; i < 4; i++ {
		fmt.Printf("%d + 2 = %d\n", i, add2(i))
		fmt.Printf("%d + 3 = %d\n", i, add3(i))
	}

	var f = acc()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	f2 := acc()
	for i := 0; i < 10; i++ {
		fmt.Println(f2(i))
	}
}
