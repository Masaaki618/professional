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
	var uniqueCount int
	arr := []string{"HND", "NRT", "KIX", "NGO", "NGO"}
	uniqueMap := make(map[string]bool)

	for _, v := range arr {
		if !uniqueMap[v] {
			uniqueMap[v] = true
		} else {
			uniqueCount++
		}
	}

	if uniqueCount == 0 {
		fmt.Fprintln(w, "false")
	} else {
		fmt.Fprintln(w, "true")
	}
}
