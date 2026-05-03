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

	for _, v := range x {
		fmt.Fprintln(w, string(v))
	}
}
