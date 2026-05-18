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
	var N, K int
	fmt.Fscan(r, &N, &K)
	pages := make([]int, N+1)
	for i := 1; i <= N; i++ {
		var x int
		fmt.Fscan(r, &x)
		pages[i] = x + pages[i-1]
	}
	pageThreshold := float64(N) / 3.0
	for i := 0; i < K; i++ {
		var a, b int
		var foulLosingA, foulLosingB bool
		var a1, a2, b1, b2 int
		fmt.Fscan(r, &a1, &a2, &b1, &b2)

		foulLosingA = isFoulPage(pageThreshold, a1, a2)
		foulLosingB = isFoulPage(pageThreshold, b1, b2)

		if foulLosingA && foulLosingB {
			fmt.Fprintln(w, "DRAW")
			continue
		} else if foulLosingB {
			fmt.Fprintln(w, "A")
			continue
		} else if foulLosingA {
			fmt.Fprintln(w, "B")
			continue
		}

		a = pages[a2] - pages[a1-1]
		b = pages[b2] - pages[b1-1]

		if a == b {
			fmt.Fprintln(w, "DRAW")
		} else if a < b {
			fmt.Fprintln(w, "B")
		} else if b < a {
			fmt.Fprintln(w, "A")
		}
	}
}

func isFoulPage(pageThreshold float64, l, r int) bool {
	return pageThreshold <= float64(r-l+1)
}
