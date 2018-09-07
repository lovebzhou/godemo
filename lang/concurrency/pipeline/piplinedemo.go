package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("gen:", n)
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			n2 := n * n
			fmt.Println("sq:", n, n2)
			out <- n2
		}
		close(out)
	}()
	return out
}

func test1() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func test2() {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func merge1(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			fmt.Println("merge1: out<-", n)
			out <- n
		}
		fmt.Println("merge1:wg.Done()")
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func test3() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge1(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func merge2(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			fmt.Println(n, "<-sq")
			select {
			case out <- n:
				fmt.Println("merge2: out<-", n)
			case <-done:
				fmt.Println("merge2: <-done")
			}
		}
		fmt.Println("merge2:wg.Done()")
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func test4() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2)
	out := merge2(done, c1, c2)
	fmt.Println("test4:", <-out, "<-out") // 4 or 9

	fmt.Println(time.Now())
	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("sleep:", i)
	}
}

// The go statement that starts a new goroutine happens before the goroutine's execution begins.
func test5() {
	var a = "test5: world"
	go func() {
		a = "test5: hello"
	}()

	fmt.Println(a)
}

func test6() {
	var done = make(chan struct{})
	var a string

	go func() {
		a = "test6: Hello, world!"
		done <- struct{}{}
	}()

	<-done

	fmt.Println(a)
}

func test7() {
	c := make(chan int)
	a := "test7: init"

	go func() {
		a = "test7: Hello world"
		<-c
		fmt.Println("1")
	}()

	fmt.Println("2")
	c <- 0
	fmt.Println("3")

	fmt.Println(a)
}

func test8() {
	l := new(sync.Mutex)

	a := "test8:"

	l.Lock()
	go func() {
		a = "test8: Hello world"
		l.Unlock()
	}()

	l.Lock()
	fmt.Println(a)
	l.Unlock()
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
	test6()
	test7()
	test8()
}
