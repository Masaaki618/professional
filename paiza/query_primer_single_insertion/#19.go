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
	result := make([]int, 100)

	for i := 0; i < 100; i++ {
		row := make([]int, 100)
		for j := 0; j < 100; j++ {
			fmt.Fscan(r, &row[j])
		}
		max := row[0]
		for _, v := range row {
			if max < v {
				max = v
			}
		}

		result[i] = max
	}
	for _, v := range result {
		fmt.Fprintln(w, v)
	}
}
