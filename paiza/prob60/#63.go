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
	var x, x1 int
	fmt.Fscan(r, &x)
	userMetadataMap := make(map[string]string)
	metadataMap := make(map[string]string)
	user := make([]string, 0, x)

	for i := 0; i < x; i++ {
		var y, z string
		fmt.Fscan(r, &y, &z)
		userMetadataMap[y] = z
		user = append(user, y)
	}
	fmt.Fscan(r, &x1)
	for i := 0; i < x1; i++ {
		var y, z string
		fmt.Fscan(r, &y, &z)
		metadataMap[y] = z
	}

	for _, v := range user {
		fmt.Fprintln(w, v, metadataMap[userMetadataMap[v]])
	}

}
