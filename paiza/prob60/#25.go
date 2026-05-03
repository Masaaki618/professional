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
	for i, v := range line {
		if x <= (i+1) && (i+1) <= y {
			fmt.Fprint(w, strings.ToUpper(string(v)))
		} else {
			fmt.Fprint(w, string(v))
		}
	}
}
