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

	y := x[0]
	z := x[len(x)-1]
	fmt.Fprintln(w, y < z)
}
