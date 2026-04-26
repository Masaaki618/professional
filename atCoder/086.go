package main

import "fmt"

func main() {
	var int1, int2 int

	fmt.Scan(&int1, &int2)
	if (int1*int2)%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
