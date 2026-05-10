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
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		var x int
		fmt.Fscan(r, &x)
		arr[i] = x
	}

	prefix := make([]int, len(arr)+1)
	for i, v := range arr {
		prefix[i+1] = prefix[i] + v
	}

	// クエリはO(1)で答えられる
	for i := 0; i < y; i++ {
		var x int
		fmt.Fscan(r, &x)
		fmt.Fprintln(w, prefix[x])
	}
}
