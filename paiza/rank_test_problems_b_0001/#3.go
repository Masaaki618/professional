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
	var x, newString string
	fmt.Fscan(r, &x)
	for _, v := range x {
		if string(v) == "t" || string(v) == "a" {
			continue
		}
		newString += string(v)
	}

	fmt.Fprintln(w, newString)
}
