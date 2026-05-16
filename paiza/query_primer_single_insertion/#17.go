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
	var H, W, N int
	fmt.Fscan(r, &H, &W, &N)
	H++
	W++
	matrix := make([][]int, H)
	matrix[0] = make([]int, W)

	for i := 1; i < H; i++ {
		row := make([]int, W)
		for j := 1; j < W; j++ {
			var x int
			fmt.Fscan(r, &x)
			row[j] = matrix[i-1][j] + row[j-1] + x - matrix[i-1][j-1]
		}
		matrix[i] = row
	}

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
	}
}
