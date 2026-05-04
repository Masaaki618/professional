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
	var n int
	fmt.Fscan(r, &n)
	uniqueKey := make(map[int]string)
	for i := 0; i < n; i++ {
		var s string
		var d int
		fmt.Fscan(r, &s, &d)
		uniqueKey[d] = s
	}
	arr := make([]int, 0, n)
	for i := range uniqueKey {
		arr = append(arr, i)
	}

	sort.Ints(arr)
	for _, v := range arr {
		fmt.Fprintln(w, uniqueKey[v])
	}
}
