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
	var diagTLBR, diagTRBL string
	for i := 0; i < 5; i++ {
		var line string
		fmt.Fscan(r, &line)
		diagTLBR += string(line[i])
		diagTRBL += string(line[4-i])
	}

	for _, diag := range [2]string{diagTLBR, diagTRBL} {
		if diag == "OOOOO" {
			fmt.Fprintln(w, "O")
			return
		} else if diag == "XXXXX" {
			fmt.Fprintln(w, "X")
			return
		}
	}

	fmt.Fprintln(w, "D")
}
