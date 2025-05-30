package main

import (
	"bytes"
	"fmt"
)

// intsToStringはfmt.Sprint(values)に似ていますがカンマを追加します。
func intsToString(values []int) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1,2,3]"
}
