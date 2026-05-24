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
	list := make([]string, H)
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(r, &s)
		list[i] = s
	}
	fmt.Fscan(r, &x, &y)
	// 文字列
	listRow := list[x]
	listByte := []byte(listRow)
	if listByte[y] == '.' {
		listByte[y] = '#'
	} else {
		listByte[y] = '.'
	}
	list[x] = string(listByte)

	for _, v := range list {
		fmt.Fprintln(w, v)

	}
}
