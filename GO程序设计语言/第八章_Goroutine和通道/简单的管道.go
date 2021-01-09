package main

import "fmt"

func main() {
	way1()
	fmt.Println("------ way2 ---------")
	way2()
	fmt.Println("------ way3 ---------")
	way3()
}

// 写法1
func way1() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		defer close(naturals)
		for x := 0; x < 5; x++ {
			naturals <- x
		}
	}()

	go func() {
		defer close(squares)
		for {
			if x, ok := <-naturals; ok {
				squares <- x * x
			} else {
				break
			}
		}
	}()

	for {
		if x, ok := <-squares; ok {
			fmt.Println(x)
		} else {
			fmt.Println("end")
			break
		}
	}
}

// 写法2
func way2() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		defer close(naturals)
		for x := 6; x < 10; x++ {
			naturals <- x
		}
	}()

	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x * x
		}

	}()

	for x := range squares {
		fmt.Println(x)
	}
	fmt.Println("end")
}

func way3() {
	// Define Function
	funcCounter := func(out chan<- int) {
		for x := 0; x < 5; x++ {
			out <- x
		}
		close(out)
	}

	funcSquarer := func(in <-chan int, out chan<- int) {
		for x := range in {
			out <- x * x
		}
		close(out)
	}

	funcPrinter := func(in <-chan int) {
		for x := range in {
			fmt.Println(x)
		}
		fmt.Println("end")
	}

	// start
	naturals := make(chan int)
	squares := make(chan int)

	go funcCounter(naturals)
	go funcSquarer(naturals, squares)
	funcPrinter(squares)
}
