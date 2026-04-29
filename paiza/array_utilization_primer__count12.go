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
	fmt.Fscan(r, &x)
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Fscan(r, &arr[i])
	}
	fmt.Fscan(r, &y)
	arr = append(arr, y)

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
