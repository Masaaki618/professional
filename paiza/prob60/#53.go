package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	resultMap := make(map[string]int)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	for _, v := range line {
		resultMap[string(v)]++
	}

	if resultMap["O"] == 5 {
		fmt.Fprintln(w, "O")
	} else if resultMap["X"] == 5 {
		fmt.Fprintln(w, "X")
	} else {
		fmt.Fprintln(w, "D")
	}

}
