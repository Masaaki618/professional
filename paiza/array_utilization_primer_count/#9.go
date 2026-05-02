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
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for i := x - 1; -1 < i; i-- {
		fmt.Fprintln(w, arr[i])
	}
}
