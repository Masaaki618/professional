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
	var x, y int
	fmt.Fscan(r, &x, &y)
	studentMap := make(map[int]string)
	for i := 0; i < x; i++ {
		var userId int
		var name string
		fmt.Fscan(r, &userId, &name)
		studentMap[userId] = name
	}

	for i := 0; i < y; i++ {
		var userId int
		fmt.Fscan(r, &userId)
		fmt.Fprintln(w, studentMap[userId])
	}
}
