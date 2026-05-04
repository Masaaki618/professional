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
	arr := []string{"HND", "NRT", "KIX", "NGO", "NGO", "NGO", "NGO", "NGO"}
	uniqueMap := make(map[string]int)

	for _, v := range arr {
		uniqueMap[v]++
	}

	for _, v := range uniqueMap {
		if v > 1 {
			fmt.Fprintln(w, v)
		}
	}
}
