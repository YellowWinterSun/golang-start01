package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1,2,3}))

	buf := bytes.NewBufferString("Hello")
	buf.WriteString(",world")
	fmt.Println(buf.String())

	buf2 := bytes.NewBufferString("黄冬")
	r, z, _ := buf2.ReadRune()
	fmt.Println("r=", string(r), ", z=", z)
	r, z, _ = buf2.ReadRune()
	fmt.Println("r=", string(r), ", z=", z)

}

func intsToString(values []int) string {
	var buf bytes.Buffer

	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	return buf.String()
}

