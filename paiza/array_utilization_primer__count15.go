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

	for i := 1; i < 10; i++ {
		fmt.Fprintln(w, 1*i, 2*i, 3*i, 4*i, 5*i, 6*i, 7*i, 8*i, 9*i)
	}
}
