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
	var n, q int
	fmt.Fscan(r, &n, &q)
	arr := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < q; i++ {
		var command string
		fmt.Fscan(r, &command)

		switch command {
		case "reverse":
			for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
				arr[i], arr[j] = arr[j], arr[i]
			}
		case "resize":
			var c int
			fmt.Fscan(r, &c)
			if len(arr) > c {
				arr = arr[:c]
			}
		case "swap":
			var a, b int
			fmt.Fscan(r, &a, &b)
			arr[a-1], arr[b-1] = arr[b-1], arr[a-1]
		}
	}

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
