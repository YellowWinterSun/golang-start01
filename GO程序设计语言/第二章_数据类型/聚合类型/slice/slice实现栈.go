package main

import "fmt"

func main() {
	stack := newStack(2)
	fmt.Println(stack, len(*stack), cap(*stack))

	stackPush(stack, 1)
	stackPush(stack, 2)
	stackPush(stack, 3)

	fmt.Println(stack)
	fmt.Println(stackPop(stack))
	fmt.Println(stack, len(*stack), cap(*stack))
	fmt.Println(stackPop(stack))
	fmt.Println(stack, len(*stack), cap(*stack))
	fmt.Println(stackPop(stack))
	fmt.Println(stack, len(*stack), cap(*stack))

	fmt.Println(stackPop(stack))
}

func newStack(cap int) (stack *[]int) {
	newArray := make([]int, 0, cap)
	stack = &newArray
	return
}

func stackPush(stack *[]int, value int) {
	*stack = append(*stack, value)
}

func stackPop(stack *[]int) (topValue int, success bool) {
	if len(*stack) < 1 {
		topValue = -1
		success = false
		return
	}

	success = true
	topValue = (*stack)[len(*stack) - 1]
	*stack = (*stack)[:len(*stack) - 1]
	return
}
