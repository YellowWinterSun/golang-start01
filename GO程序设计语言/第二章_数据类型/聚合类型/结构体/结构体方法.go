package main

import "fmt"

type MyPoint struct {
	X, Y int
}

type MyName struct {
	Name string
}

type MyColor struct {
	A *MyPoint
	*MyName
}

func (p *MyPoint) MyPointSay() {
	fmt.Println("MyPointSay:", p.X, p.Y)
}

func (p *MyName) MyNameSay() {
	fmt.Println("MyNameSay:", p.Name)
}

func main() {
	c1 := MyColor{&MyPoint{1,2}, &MyName{"DG"}}
	c1.MyNameSay()
	c1.A.MyPointSay()

	funcP := c1.MyNameSay
	funcP()	// 相当于调用c1.MyNameSay()
}
