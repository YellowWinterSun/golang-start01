package main

import (
	"fmt"
	"sort"
)

/**
实现这3个方法（sort.Interface），即可对数组元素进行排序
*/
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/**
自定义结构体
*/
type Track struct {
	Title  string
	Artist string
	Rank   int8
}

// 根据Artist字段进行排序
type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// 根据Rank排序
type byRank []*Track

func (x byRank) Len() int {
	return len(x)
}

func (x byRank) Less(i, j int) bool {
	return x[i].Rank < x[j].Rank
}

func (x byRank) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	var array StringSlice = []string{"a", "c", "b", "z"}

	sort.Sort(array)
	fmt.Println(array)

	fmt.Println("------- 自定义结构体排序 ---------")
	fmt.Println(">>> 根据Artist升序排序")
	var trackArray = []*Track{
		{"title1", "A", 3},
		{"title2", "C", 2},
		{"title3", "B", 1},
		{"title4", "D", 4},
	}

	sort.Sort(byArtist(trackArray))
	for _, v := range trackArray {
		fmt.Printf("%v\n", v)
	}

	// 降序排序，不需要我们重新定义一个新的type，然后改写Less的判断逻辑。只需要用sort.Reverse即可
	fmt.Println(">>> 根据Rank降序排序")
	sort.Sort(sort.Reverse(byRank(trackArray)))
	for _, v := range trackArray {
		fmt.Printf("%v\n", v)
	}
}
