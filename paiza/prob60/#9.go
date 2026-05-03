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
	var x int
	fmt.Fscan(r, &x)
	r.ReadString('\n')

	for i := 0; i < x; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		fmt.Fprintln(w, line)
	}
}
