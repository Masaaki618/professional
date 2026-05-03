package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var x, y string
	fmt.Fscan(r, &x, &y)

	if n, err := strconv.Atoi(x); err != nil {
		fmt.Fprintln(w, x)
	} else {
		fmt.Fprintln(w, n)
	}

	if n, err := strconv.Atoi(y); err != nil {
		fmt.Fprintln(w, y)
	} else {
		fmt.Fprintln(w, n)
	}
}
