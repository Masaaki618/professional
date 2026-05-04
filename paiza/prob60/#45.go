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
		var y string
		var z int
		fmt.Fscan(r, &y, &z)

		fmt.Fprintln(w, z)
	}
}
