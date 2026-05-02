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

	targetPin := 0
	standingCount := 0

	for i := 0; i < 10; i++ {
		var x int
		fmt.Fscan(r, &x)
		if x == 1 {
			targetPin = 10 - i
			standingCount++
		}
	}

	fmt.Fprintln(w, targetPin)
	fmt.Fprintln(w, standingCount)
}
