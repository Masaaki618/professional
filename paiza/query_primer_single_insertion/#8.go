package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var x, y, z int
	fmt.Fscan(r, &x, &y, &z)
	arr := make([]int, 0, x+2)
	arr = append(arr, y, z)
	for i := 0; i < x; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}
	sort.Ints(arr)

	for i, v := range arr {
		if v == z {
			fmt.Fprintln(w, i+1)
		}
	}
}
