package main

import (
	"fmt"
	"math"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("ok")
}

func TestB(t *testing.T) {
	fmt.Println("ok-B")
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := math.Pow(float64(2), float64(5))
		x *= x
	}
}
