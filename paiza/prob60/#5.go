package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var value int
	arr := []int{4, 0, 5, -1, 3, 10, 6, -8}

	for _, v := range arr {

		if 5 <= v {
			value += v
		}
	}

	fmt.Fprintln(w, value)
}
