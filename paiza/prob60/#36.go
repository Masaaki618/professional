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
	var y string
	var z int
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)

	for i := 0; i+len(x) <= len(y); i++ {
		if y[i:i+len(x)] == x {
			z++
		}
	}

	fmt.Fprintln(w, z)
}
