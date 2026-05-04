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

	sumMap := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		var d int
		fmt.Fscan(r, &s, &d)
		sumMap[s] += d
	}

	keys := make([]string, 0, len(sumMap))
	for k := range sumMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return sumMap[keys[i]] > sumMap[keys[j]]
	})

	for _, k := range keys {
		fmt.Fprintln(w, k, sumMap[k])
	}
}
