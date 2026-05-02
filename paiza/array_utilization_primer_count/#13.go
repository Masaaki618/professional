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
	fmt.Fscan(r, &x)
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Fscan(r, &arr[i])
	}

	fmt.Fscan(r, &y, &z)
	for i, _ := range arr {
		if (1 + i) == y {
			arr = append(arr[:i+1], append([]int{z}, arr[i+1:]...)...)
		}
	}

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
