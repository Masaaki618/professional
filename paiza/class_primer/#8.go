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

type Point struct {
	alpha string
	road1 int
	road2 int
}

func main() {
	defer w.Flush()
	var N, K, S int
	fmt.Fscan(r, &N, &K, &S)

	points := make(map[int]Point, N)
	for i := 1; i <= N; i++ {
		var alpha string
		var road1, road2 int
		fmt.Fscan(r, &alpha, &road1, &road2)
		points[i] = Point{alpha, road1, road2}
	}

	current := S
	fmt.Fprintf(w, points[current].alpha)

	for i := 0; i < K; i++ {
		var m int
		fmt.Fscan(r, &m)
		if m == 1 {
			current = points[current].road1
		} else {
			current = points[current].road2
		}
		fmt.Fprintf(w, points[current].alpha)
	}
	fmt.Fprintln(w)
}
