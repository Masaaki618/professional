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
	arr := make([]int, y)
	for i := 0; i < y; i++ {
		fmt.Fscan(r, &arr[i])
	}

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}

}
