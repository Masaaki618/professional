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
	var x int
	fmt.Fscan(r, &x)
	weight := make([]int, 5)

	for i := 0; i < len(weight); i++ {
		fmt.Fscan(r, &weight[i])
	}

	maxValue := 0
	// 生徒の人数分
	for i := 0; i < x; i++ {
		totalValue := 0
		for j := 0; j < len(weight); j++ {
			var value int
			fmt.Fscan(r, &value)
			totalValue += weight[j] * value
		}

		if maxValue < totalValue {
			maxValue = totalValue
		}
	}

	fmt.Fprintln(w, maxValue)
}
