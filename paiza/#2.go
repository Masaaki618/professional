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
	var x, y, z int
	fmt.Fscan(r, &x, &y, &z)
	fmt.Fprintln(w, x+y+z)
	fmt.Fprintln(w, x*y*z)
}
