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
		var nickname, birth, state string
		var old int
		fmt.Fscan(r, &nickname, &old, &birth, &state)
		arr = append(arr, User{
			nickname: nickname,
			old:      old,
			birth:    birth,
			state:    state,
		})
	}
	for _, v := range arr {
		s := fmt.Sprintf(`User{
nickname : %s
old : %d
birth : %s
state : %s
}`, v.nickname, v.old, v.birth, v.state)
		fmt.Fprintln(w, s)
	}

}
