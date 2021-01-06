package main

import "fmt"

func main() {
	a := [...]int{1, 2 ,3}
	zero1(a)
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	zero2(&b)
	fmt.Println(b)
}

func zero1(a [3]int) {
	a[0] = 0
}

func zero2(ptr *[3]int) {
	(*ptr)[0] = 0
}
