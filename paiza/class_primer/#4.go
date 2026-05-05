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

type User struct {
	nickname string
	old      int
	birth    string
	state    string
}

func main() {
	defer w.Flush()
	var x, y int
	fmt.Fscan(r, &x, &y)

	userMap := make(map[int]*User)

	for i := 1; i <= x; i++ {
		var u User
		fmt.Fscan(r, &u.nickname, &u.old, &u.birth, &u.state)
		userMap[i] = &u
	}

	for i := 1; i <= y; i++ {
		var userId int
		var name string
		fmt.Fscan(r, &userId, &name)
		user := userMap[userId]
		user.changeName(name)
	}

	for i := 1; i <= len(userMap); i++ {
		u := userMap[i]
		fmt.Fprintln(w, u.nickname, u.old, u.birth, u.state)
	}
}

func (u *User) changeName(name string) {
	u.nickname = name
}
