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

	result := "false"
	if x[0] <= z[0] && z[0] <= y[0] {
		result = "true"
	}

	fmt.Fprintln(w, result)
}
