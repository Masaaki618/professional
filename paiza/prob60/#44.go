package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 5, 6, 3, 2, 5, 23, 2}
	sort.Ints(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
