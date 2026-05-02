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
	var n, k, m int
	fmt.Fscan(r, &n, &k, &m)

	count := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		if a >= k {
			count++
		}
	}

	result := count - m
	if result < 0 {
		result = 0
	}
	fmt.Fprintln(w, result)
}
