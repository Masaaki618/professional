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
	arr := make([]int, 0, x)
	for i := 0; i < x; i++ {
		var z int
		fmt.Fscan(r, &z)
		if y <= z {
			arr = append(arr, z)
		}
	}

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
