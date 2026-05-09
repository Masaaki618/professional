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

func main() {
	defer w.Flush()
	var x, y, z int
	fmt.Fscan(r, &x, &y, &z)
	arr := make([]int, 0, x+y+2)
	arr = append(arr, z)
	for i := 0; i < x; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}
	sort.Ints(arr) // 最初に1回だけソート

	for i := 0; i < y; i++ {
		var medhods string
		fmt.Fscan(r, &medhods)
		switch medhods {
		case "join":
			var x int
			fmt.Fscan(r, &x)
			pos := sort.SearchInts(arr, x) // 挿入位置をバイナリサーチ
			arr = append(arr, 0)
			copy(arr[pos+1:], arr[pos:]) // 後ろにずらす
			arr[pos] = x                 // 正しい位置に挿入
		case "sorting":
			pos := sort.SearchInts(arr, z) // zの位置をバイナリサーチ
			fmt.Fprintln(w, pos+1)
		}
	}
}
