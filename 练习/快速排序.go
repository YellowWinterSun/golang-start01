package main

import "fmt"

func main() {
	array := []int{8,5,3,2,4,1,10,7,6}
	sort(array)

	fmt.Println(array)
}

func sort(array []int) {
	if array == nil {
		return
	}
	if len(array) < 2 {
		return
	}

	quickSort(array)
}

// go语言可以利用切片技术
func quickSort(array []int) {
	p := 0
	q := len(array) - 1
	mid := array[p]

	for p < q {
		for p < q && array[q] > mid {
			q--
		}
		if p < q {
			swap(array, p, q)
		}

		for p < q && array[p] <= mid {
			p++
		}
		if p < q {
			swap(array, p, q)
		}
	}

	if p > 1 {
		quickSort(array[:p])
	}
	if q < len(array) - 1 {
		quickSort(array[q + 1:])
	}
}

func swap(array []int, left int, right int) {
	array[left], array[right] = array[right], array[left]
}
