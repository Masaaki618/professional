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

	fmt.Fscan(r, &y)
	for _, v := range arr {
		if v.old == y {
			fmt.Fprintln(w, v.nickname)
		}
	}

}
