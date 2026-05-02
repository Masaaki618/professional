package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	sort.Ints(arr)

	minDiff := arr[1] - arr[0]
	lo, hi := arr[0], arr[1]

	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
			lo, hi = arr[i-1], arr[i]
		}
	}

	fmt.Fprintln(w, lo)
	fmt.Fprintln(w, hi)
}
