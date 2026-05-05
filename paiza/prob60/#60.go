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
	var targetName string
	fmt.Fscan(r, &targetName)
	var x int
	fmt.Fscan(r, &x)

	for i := 0; i < x; i++ {
		var y, z string
		fmt.Fscan(r, &y, &z)
		if targetName == y {
			fmt.Fprintln(w, y, z)
		}
	}
}
