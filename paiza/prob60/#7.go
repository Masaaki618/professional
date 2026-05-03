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
	var x, totalCount int
	fmt.Fscan(r, &x)
	for i := 0; i < x; i++ {
		var value int
		fmt.Fscan(r, &value)
		if 5 <= value {
			totalCount += value
		}
	}

	fmt.Fprintln(w, totalCount)
}
