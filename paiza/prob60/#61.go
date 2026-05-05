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
	var x string
	fmt.Fscan(r, &x)
	var y int
	fmt.Fscan(r, &y)

	for i := 0; i < y; i++ {
		var a, b string
		fmt.Fscan(r, &a, &b)

		if x == a {
			fmt.Fprintln(w, b)
		}
	}
}
