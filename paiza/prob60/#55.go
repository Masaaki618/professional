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
	cols := [5]string{}
	for i := 0; i < 5; i++ {
		var line string
		fmt.Fscan(r, &line)
		for j := 0; j < 5; j++ {
			cols[j] += string(line[j])
		}
	}

	for _, col := range cols {
		if col == "OOOOO" {
			fmt.Fprintln(w, "O")
			return
		} else if col == "XXXXX" {
			fmt.Fprintln(w, "X")
			return
		}
	}

	fmt.Fprintln(w, "D")
}
