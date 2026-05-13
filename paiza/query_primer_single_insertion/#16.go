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

	for i := 0; i < y; i++ {
		var x, y, total int
		fmt.Fscan(r, &x, &y)
		for _, v := range arr[x-1 : y] {
			total += v
		}
		fmt.Fprintln(w, total)
	}
}
