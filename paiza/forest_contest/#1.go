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
	var x string
	fmt.Fscan(r, &x)

	sportsMap := map[string]int{
		"baseball":   9,
		"soccer":     11,
		"rugby":      15,
		"basketball": 5,
		"volleyball": 6,
	}

	fmt.Fprintln(w, sportsMap[x])
}
