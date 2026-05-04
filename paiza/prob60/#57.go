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
	result := make([]string, 5)
	var diagTLBR, diagTRBL string

	for i := 0; i < 5; i++ {
		var line string
		fmt.Fscan(r, &line)
		// 横の判定
		if line == "XXXXX" {
			fmt.Fprintln(w, "X")
			return
		} else if line == "OOOOO" {
			fmt.Fprintln(w, "O")
			return
		}

		// 縦の判定
		for j := 0; j < len(line); j++ {
			x := string(line[j])
			result[j] += x
		}

		// 斜めの判定
		diagTLBR += string(line[i])
		diagTRBL += string(line[4-i])
	}

	result = append(result, diagTLBR, diagTRBL)
	for _, v := range result {
		if v == "XXXXX" {
			fmt.Fprintln(w, "X")
			return
		} else if v == "OOOOO" {
			fmt.Fprintln(w, "O")
			return
		}
	}

	fmt.Fprintln(w, "D")
}
