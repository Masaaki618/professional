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
	result := make([]int, x+1)
	for i := 1; i <= x; i++ {
		var num int
		fmt.Fscan(r, &num)
		result[i] = result[i-1] + num
	}

	for i := 0; i < y; i++ {
		var num int
		fmt.Fscan(r, &num)
		fmt.Fprintln(w, result[num])
	}
}
