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

	var n int
	fmt.Fscan(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprintln(w, arr[j]*arr[i])
		}
	}

}
