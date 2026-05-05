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
	var targetUserName, userMetadata string
	var x, x1 int
	fmt.Fscan(r, &targetUserName)
	fmt.Fscan(r, &x)
	for i := 0; i < x; i++ {
		var y, z string
		fmt.Fscan(r, &y, &z)

		if y == targetUserName {
			userMetadata = z
		}
	}

	fmt.Fscan(r, &x1)
	for i := 0; i < x1; i++ {
		var y, z string
		fmt.Fscan(r, &y, &z)
		if y == userMetadata {
			fmt.Fprintln(w, z)
			break
		}
	}
}
