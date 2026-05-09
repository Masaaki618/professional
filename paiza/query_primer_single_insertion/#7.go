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
		var studentId int
		var name string
		fmt.Fscan(r, &studentId, &name)
		studentMap[studentId] = name
	}

	for i := 0; i < y; i++ {
		var methods string
		var studentId int
		fmt.Fscan(r, &methods, &studentId)
		switch methods {
		case "call":
			fmt.Fprintln(w, studentMap[studentId])
		case "leave":
			delete(studentMap, studentId)
		case "join":
			var name string
			fmt.Fscan(r, &name)
			studentMap[studentId] = name
		}
	}
}
