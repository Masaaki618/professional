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

	if x <= 6 {
		fmt.Fprintln(w, "no break")
	} else if x <= 8 {
		fmt.Fprintln(w, "45 min")
	} else {
		fmt.Fprintln(w, "60 min")
	}
}
