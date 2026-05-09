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
	var x, y int
	fmt.Fscan(r, &x, &y)
	var userList []string
	for i := 0; i < x; i++ {
		var user string
		fmt.Fscan(r, &user)
		userList = append(userList, user)
	}
	sort.Strings(userList)

	for i := 0; i < y; i++ {
		var medhods string
		fmt.Fscan(r, &medhods)
		switch medhods {
		case "handshake":
			for _, v := range userList {
				fmt.Fprintln(w, v)
			}
		case "leave":
			var name string
			fmt.Fscan(r, &name)
			pos := sort.SearchStrings(userList, name)
			if pos < len(userList) && userList[pos] == name {
				userList = append(userList[:pos], userList[pos+1:]...)
			}
		case "join":
			var name string
			fmt.Fscan(r, &name)
			pos := sort.SearchStrings(userList, name)
			userList = append(userList, "")
			copy(userList[pos+1:], userList[pos:])
			userList[pos] = name
		}
	}
}
