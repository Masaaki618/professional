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
	var H, W, y, x int
	fmt.Fscan(r, &H, &W)
	board := make([][]byte, H)
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(r, &s)
		board[i] = []byte(s)
	}
	fmt.Fscan(r, &y, &x)
	if board[y][x] == '.' {
		board[y][x] = '#'
	} else {
		board[y][x] = '.'
	}
	for _, row := range board {
		fmt.Fprintln(w, string(row))
	}
}
