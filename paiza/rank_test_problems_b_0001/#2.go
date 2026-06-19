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
	var y string
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)
	for i := 0; i < x; i++ {
		fmt.Fprintln(w, y)
	}
}
