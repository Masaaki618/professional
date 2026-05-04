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
	for i := 0; i < 5; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		fmt.Fprintln(w, line)
	}
}
