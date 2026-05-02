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
	var x int
	fmt.Fscan(r, &x)

	uniqueKeyMap := make(map[int]bool)
	result := make([]int, 0, x)

	for i := 0; i < x; i++ {
		var n int
		fmt.Fscan(r, &n)
		if !uniqueKeyMap[n] {
			uniqueKeyMap[n] = true
			result = append(result, n)
		}
	}

	for _, v := range result {
		fmt.Fprintln(w, v)

	}
}
