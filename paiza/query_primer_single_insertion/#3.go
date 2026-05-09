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
	var x, y int
	fmt.Fscan(r, &x, &y)
	uniqueMap := make(map[int]bool)
	for i := 0; i < x; i++ {
		var z int
		fmt.Fscan(r, &z)
		uniqueMap[z] = true
	}

	for i := 0; i < y; i++ {
		var z int
		fmt.Fscan(r, &z)
		if uniqueMap[z] {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
