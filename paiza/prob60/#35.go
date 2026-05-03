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
	var n int
	var s string
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &s)

	if n+1 <= len(s) {
		fmt.Fprintf(w, "%s %s\n", s[n-1:n], s[n:n+1])
	}
}
