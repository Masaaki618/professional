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
	var count int
	fmt.Fscan(r, &x)
	fmt.Fscan(r, &y)

	for _, v := range y {
		if string(v) == x {
			count++
		}
	}

	fmt.Fprintln(w, count)

}
