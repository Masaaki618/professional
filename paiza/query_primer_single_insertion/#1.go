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
	var x, y, z int
	fmt.Fscan(r, &x, &y, &z)
	arr := make([]int, 0, x+1)
	for i := 0; i < x; i++ {
		var x int
		fmt.Fscan(r, &x)
		arr = append(arr, x)
	}

	arr = append(arr[:y], append([]int{z}, arr[y:]...)...)
	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
