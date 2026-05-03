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

	if 5 <= x {
		fmt.Fprintln(w, "high")
	} else {
		fmt.Fprintln(w, "low")
	}
}
