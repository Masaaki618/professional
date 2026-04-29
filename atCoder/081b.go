package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var count int
	for allEven(nums) {
		for i := range nums {
			nums[i] /= 2
		}
		count++
	}

	fmt.Println(count)
}

func allEven(nums []int) bool {
	for _, v := range nums {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
