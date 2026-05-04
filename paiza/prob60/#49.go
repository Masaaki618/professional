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

	start := x[0]
	end := x[len(x)-1]

	for c := start; c <= end; c++ {
		fmt.Fprintln(w, string(c))
	}
}
