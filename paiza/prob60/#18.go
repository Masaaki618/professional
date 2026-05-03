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
	var x, y int
	fmt.Fscan(r, &x, &y)
	r.ReadString('\n')
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	fmt.Fprintln(w, line[x-1:y])
}
