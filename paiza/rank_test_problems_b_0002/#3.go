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
	var N, Q, aCount int
	fmt.Fscan(r, &N, &Q)
	m := make(map[int]string, N)
	for i := 0; i < Q; i++ {
		var p string
		var x int
		fmt.Fscan(r, &p, &x)
		if _, ok := m[x]; !ok {
			m[x] = p
		}
	}
	for _, v := range m {
		if v == "A" {
			aCount++
		}
	}
	bCount := len(m) - aCount

	if aCount == bCount {
		fmt.Fprintln(w, "Draw")
	} else if aCount > bCount {
		fmt.Fprintln(w, "A")
	} else {
		fmt.Fprintln(w, "B")
	}
}
