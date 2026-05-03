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
	arr := []string{"Nara", "Shiga", "Hokkaido", "Chiba"}
	for _, v := range arr {
		fmt.Fprintln(w, v)
	}
}
