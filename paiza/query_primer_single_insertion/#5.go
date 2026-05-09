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
	var x, y int
	fmt.Fscan(r, &x, &y)
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		var z int
		fmt.Fscan(r, &z)
		arr[i] = z
	}

	for i := 0; i < y; i++ {
		var str string
		fmt.Fscan(r, &str)
		switch str {
		case "pop":
			if len(arr) > 1 {
				arr = arr[1:]
			}
		case "show":
			for _, v := range arr {
				fmt.Fprintln(w, v)
			}
		}
	}
}
