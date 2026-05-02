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

	if y >= 1 && y <= len(arr) {
		arr = append(arr[:y-1], arr[y:]...)
	}

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
