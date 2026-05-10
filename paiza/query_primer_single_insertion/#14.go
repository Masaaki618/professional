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
	superchatUserMap := make(map[string]int)
	var superchatUserList []string
	var membershipList []string

	for i := 0; i < x; i++ {
		var name, method string
		fmt.Fscan(r, &name, &method)
		switch method {
		case "give":
			var amount int
			var dummy string
			fmt.Fscan(r, &amount, &dummy)
			superchatUserMap[name] += amount
		case "join":
			var dummy string
			fmt.Fscan(r, &dummy)
			membershipList = append(membershipList, name)
		}
	}
	for k := range superchatUserMap {
		superchatUserList = append(superchatUserList, k)
	}

	sort.Strings(membershipList)
	sort.Slice(superchatUserList, func(i, j int) bool {
		if superchatUserMap[superchatUserList[i]] != superchatUserMap[superchatUserList[j]] {
			return superchatUserMap[superchatUserList[i]] > superchatUserMap[superchatUserList[j]]
		}
		return superchatUserList[i] > superchatUserList[j]
	})

	for _, v := range append(superchatUserList, membershipList...) {
		fmt.Fprintln(w, v)
	}
}
