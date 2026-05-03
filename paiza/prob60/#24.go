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
	var x, y int
	fmt.Fscan(r, &x, &y)

	for i := x; i <= y; i++ {
		fmt.Fprintln(w, i)
	}
}
