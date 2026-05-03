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
	var x, y string
	fmt.Fscan(r, &x, &y)
	fmt.Fprintln(w, x)
	fmt.Fprintln(w, y)
}
