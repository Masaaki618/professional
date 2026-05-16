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
	W++
	H++
	matrix := make([][]int, H)
	matrix[0] = make([]int, W)

	for i := 1; i < H; i++ {
		row := make([]int, W)
		for j := 1; j < W; j++ {
			var z int
			fmt.Fscan(r, &z)
			row[j] = matrix[i-1][j] + row[j-1] + z - matrix[i-1][j-1]
		}
		matrix[i] = row
	}

	for i := 0; i < N; i++ {
		var a, b, c, d int
		fmt.Fscan(r, &a, &b, &c, &d)
		fmt.Fprintln(w, matrix[c][d]-matrix[a-1][d]-matrix[c][b-1]+matrix[a-1][b-1])
	}
}
