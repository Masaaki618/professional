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
	var n, k, f int
	fmt.Fscan(r, &n, &k, &f)
	arr := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &arr[i])
	}

	seen := make(map[int]bool)
	result := make([]int, 0, n)
	for _, v := range arr[f:] {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
