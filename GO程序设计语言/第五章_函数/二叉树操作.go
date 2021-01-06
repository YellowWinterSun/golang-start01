package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

// 前序遍历
func PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Value)

	PreOrder(node.Left)
	PreOrder(node.Right)
}

// 层数
func Layers(node *Node) int {
	if node == nil {
		return 0
	}

	leftLayers := Layers(node.Left)
	rightLayers := Layers(node.Right)

	if leftLayers >= rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

/*
      1
  2       3
4   5
 */
func main() {
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}

	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}

	PreOrder(root)
	fmt.Println()
	fmt.Println(Layers(root))
}
