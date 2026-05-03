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
	var x string
	fmt.Fscan(r, &x)
	fmt.Fprintln(w, strings.ToUpper(x))
}
