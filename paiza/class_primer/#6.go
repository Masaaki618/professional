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

type employee struct {
	number int
	name   string
}

func NewEmployee(number int, name string) employee {
	return employee{number: number, name: name}
}

func (e employee) getname() string {
	return e.name
}

func (e employee) getnum() int {
	return e.number
}

func main() {
	defer w.Flush()
	var x int
	employeeCount := 0

	fmt.Fscan(r, &x)
	employeeMap := make(map[int]employee)

	for i := 1; i <= x; i++ {
		var y string
		fmt.Fscan(r, &y)

		switch y {
		case "make":
			var number int
			var name string
			fmt.Fscan(r, &number, &name)
			employeeCount++
			employeeMap[employeeCount] = NewEmployee(number, name)
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
