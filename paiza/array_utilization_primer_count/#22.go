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
	var n int
	fmt.Fscan(r, &n)
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var cmd string
		fmt.Fscan(r, &cmd)
		if cmd == "in" {
			var val int
			fmt.Fscan(r, &val)
			queue = append(queue, val)
		} else if cmd == "out" && len(queue) > 0 {
			queue = queue[1:]
		}
	}

	for _, v := range queue {
		fmt.Fprintln(w, v)
	}
}
