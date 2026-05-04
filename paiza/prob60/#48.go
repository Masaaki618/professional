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
	for i := 0; i < 3; i++ {
		var x string
		fmt.Fscan(r, &x)
		fmt.Fprintln(w, x)
	}
}
