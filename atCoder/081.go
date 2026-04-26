package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	var count int
	fmt.Scan(&x)
	s := strconv.Itoa(x)
	for _, v := range s {
		if v == '1' {
			count++
		}
	}

	fmt.Println(count)
}
