package main

import "fmt"

func main() {
	var n, target int
	fmt.Scan(&n, &target)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	fmt.Println(count)
}
