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
		var y, z int
		fmt.Fscan(r, &y, &z)
		if y == z {
			totalCount += y * z
		} else {
			totalCount += y + z
		}
	}
	fmt.Fprintln(w, totalCount)
}
