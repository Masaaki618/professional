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
	var x int
	fmt.Fscan(r, &x)
	arr := make([]int, 0, x)

	for i := 0; i < x; i++ {
		var x int
		fmt.Fscan(r, &x)
		arr = append(arr, x)
	}

	for _, v := range arr[1:] {
		fmt.Fprintln(w, v)
	}
}
