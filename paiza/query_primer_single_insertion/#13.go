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

type order struct {
	orderId string
	amount  int
}

func NewOrder(orderId string, amount int) order {
	return order{
		orderId: orderId,
		amount:  amount,
	}
}

func main() {
	defer w.Flush()
	var x, y int
	fmt.Fscan(r, &x, &y)
	departmentList := make([]string, x)
	departmentMap := make(map[string][]order)

	for i := 0; i < x; i++ {
		var department string
		fmt.Fscan(r, &department)
		departmentList[i] = department
	}

	for i := 0; i < y; i++ {
		var department, orderId string
		var amount int
		fmt.Fscan(r, &department, &orderId, &amount)
		o := NewOrder(orderId, amount)
		departmentMap[department] = append(departmentMap[department], o)
	}

	for _, department := range departmentList {
		fmt.Fprintln(w, department)
		for _, v := range departmentMap[department] {
			fmt.Fprintln(w, v.orderId, v.amount)
		}
		fmt.Fprintln(w, "-----")
	}
}
