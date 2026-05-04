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

	for i := 0; i < x; i++ {
		var x string
		var y int
		fmt.Fscan(r, &x, &y)
		fmt.Fprintln(w, y)
	}
}
