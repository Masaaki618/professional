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
		var x int
		fmt.Fscan(r, &x)
		arr = append(arr, x)
	}

	result := "NO"
	for _, v := range arr {
		if v == y {
			result = "YES"
		}
	}

	fmt.Fprintln(w, result)

}
