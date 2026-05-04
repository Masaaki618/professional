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
	for i := 0; i < 5; i++ {
		var line string
		fmt.Fscan(r, &line)

		if line == "OOOOO" {
			fmt.Fprintln(w, "O")
			return
		} else if line == "XXXXX" {
			fmt.Fprintln(w, "X")
			return
		}
	}
	fmt.Fprintln(w, "D")
}
