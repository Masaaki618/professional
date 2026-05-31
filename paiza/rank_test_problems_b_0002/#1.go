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
	var x, y, z string
	fmt.Fscan(r, &x, &y, &z)
	fmt.Fprintln(w, x)
	fmt.Fprintln(w, y)
	fmt.Fprintln(w, z)
}
