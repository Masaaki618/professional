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
	var x int
	fmt.Fscan(r, &x)
	a := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		fmt.Fscan(r, &a[i])
	}
	for i := 0; i < x; i++ {
		var c, d, max int
		fmt.Fscan(r, &c, &d)
		max = a[c]
		for j := c; j <= d; j++ {
			if max < a[j] {
				max = a[j]
			}
		}
		fmt.Fprintln(w, max)
	}

}
