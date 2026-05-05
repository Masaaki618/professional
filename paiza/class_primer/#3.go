package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var x int
	fmt.Fscan(r, &x)

	type User struct {
		nickname string
		old      int
		birth    string
		state    string
	}

	arr := []User{}

	for i := 0; i < x; i++ {
		var u User
		fmt.Fscan(r, &u.nickname, &u.old, &u.birth, &u.state)
		arr = append(arr, u)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].old < arr[j].old
	})

	for _, v := range arr {
		fmt.Fprintln(w, v.nickname, v.old, v.birth, v.state)
	}
}
