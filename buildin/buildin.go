package main

import (
	"fmt"
	"math"
)

func makeDemo() {
	fmt.Println("============= makeDemo ===============")

	dumpSlice := func(n string, vs []int) {
		fmt.Printf("[%s (len = %d, cap = %d ) \n %v\n", n, len(vs), cap(vs), vs)
	}

	s1 := make([]int, 10, 100) // slice with len(s) == 10, cap(s) == 100
	s1[1] = 1
	s1[2] = 2
	dumpSlice("s1", s1)
	s1 = append(s1, 11)
	dumpSlice("s1", s1)

	s2 := make([]int, 1e3) // slice with len(s) == cap(s) == 1000
	s2[9] = 999
	// dumpSlice("s2", s2)

	// s3 := make([]int, 1<<63)        // illegal: len(s) is not representable by a value of type int
	// s4 := make([]int, 10, 0)        // illegal: len(s) > cap(s)

	c := make(chan int, 10) // channel with a buffer size of 10
	c <- 1
	fmt.Printf("<-c = %d\n", <-c)

	m := make(map[string]int, 100) // map with initial space for approximately 100 elements
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	fmt.Printf("m(len = %d) = %v\n", len(m), m)
}

// Appending to and copying slices
func appendDemo() {
	fmt.Println("============= appendDemo ===============")
	s0 := []int{0, 0}
	s1 := append(s0, 2)              // append a single element     s1 == []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7)        // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)          // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	s4 := append(s3[3:6], s3[2:]...) // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}
	fmt.Println(s4)

	var t []interface{}
	t = append(t, 42, 3.1415, "foo") //                             t == []interface{}{42, 3.1415, "foo"}
	fmt.Println(t)

	var b []byte
	// As a special case, append also accepts a first argument assignable to
	// type []byte with a second argument of string type followed by ...
	// This form appends the bytes of the string.
	b = append(b, "bar"...) // append string contents      b == []byte{'b', 'a', 'r' }
	fmt.Printf("b[1] = %c\n", b[1])
	fmt.Println(b)
}

// Appending to and copying slices
func copyDemo() {
	fmt.Println("============= copyDemo ===============")

	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s = make([]int, 6)
	var b = make([]byte, 5)

	n1 := copy(s, a[0:]) // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("n1 = %d, s = %v\n", n1, s)

	n2 := copy(s, s[2:]) // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	fmt.Printf("n2 = %d, s = %v\n", n2, s)

	n3 := copy(b, "Hello, World!") // n3 == 5, b == []byte("Hello")
	fmt.Printf("n3 = %d, b = %v\n", n3, b)

}

// Deletion of map elements
func deleteMap() {
	fmt.Println("============= deleteMap ===============")
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println(m1)
	delete(m1, "c")
	fmt.Println(m1)
}

func complexDemo() {
	fmt.Println("============= complexDemo ===============")

	var a = complex(2, -2)              // complex128
	const b = complex(1.0, -1.4)        // untyped complex constant 1 - 1.4i
	x := float32(math.Cos(math.Pi / 2)) // float32
	var c64 = complex(5, -x)            // complex64
	var s uint = complex(1, 0)          // untyped complex constant 1 + 0i can be converted to uint
	// _ = complex(1, 2<<s)                // illegal: 2 assumes floating-point type, cannot shift
	var rl = real(c64)                  // float32
	var im = imag(a)                    // float64
	const c = imag(b)                   // untyped constant -1.4
	// _ = imag(3 << s)                    // illegal: 3 assumes complex type, cannot shift
	fmt.Printf("a=%v\nb=%v\nx=%v\nc64=%v\ns=%v\nrl=%v\nim=%v\nc=%v\n", a,b,x,c64,s,rl,im,c)
}

func main() {
	makeDemo()
	appendDemo()
	copyDemo()
	deleteMap()
	complexDemo()
}
