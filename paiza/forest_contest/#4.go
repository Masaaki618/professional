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

type Item struct {
	key int
	val string
}

func main() {
	defer w.Flush()
	var x int
	fmt.Fscan(r, &x)
	items := make([]Item, x)

	for i := 0; i < x; i++ {
		fmt.Fscan(r, &items[i].key, &items[i].val)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].key < items[j].key
	})

	for _, item := range items {
		fmt.Fprintln(w, item.key, item.val)
	}
}
