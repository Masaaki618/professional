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
	arr := []int{1, 3, 5, 6, 3, 2, 5, 23, 2}

	for _, v := range arr {
		value += v
	}

	fmt.Fprintln(w, value)
}
