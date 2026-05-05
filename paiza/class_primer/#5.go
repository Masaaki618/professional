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

type Employee struct {
	number int
	name   string
}

func (e Employee) getname() string {
	return e.name
}

func (e Employee) getnum() int {
	return e.number
}

func main() {
	defer w.Flush()
	var x int
	EmployeeCount := 0

	fmt.Fscan(r, &x)
	employeeMap := make(map[int]Employee)

	for i := 1; i <= x; i++ {
		var y string
		fmt.Fscan(r, &y)

		switch y {
		case "make":
			var e Employee
			fmt.Fscan(r, &e.number, &e.name)
			EmployeeCount++
			employeeMap[EmployeeCount] = e
		case "getname":
			var z int
			fmt.Fscan(r, &z)
			user := employeeMap[z]
			fmt.Fprintln(w, user.getname())
		case "getnum":
			var z int
			fmt.Fscan(r, &z)
			user := employeeMap[z]
			fmt.Fprintln(w, user.getnum())
		}
	}

}
