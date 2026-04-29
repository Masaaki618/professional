package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(reader, &t)
		sum += t
	}

	fmt.Fprintln(writer, sum)
}
