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

type orderUser struct {
	age               int
	totalPrice        int
	hasOrderedAlcohol bool
	isCheckedOut      bool
}

func NewOrderUser(age int) *orderUser {
	return &orderUser{
		age:          age,
		isCheckedOut: false,
	}
}

func main() {
	defer w.Flush()
	var x, y int
	fmt.Fscan(r, &x, &y)
	userList := make([]*orderUser, 0, x)

	for i := 0; i < x; i++ {
		var age int
		fmt.Fscan(r, &age)
		userList = append(userList, NewOrderUser(age))
	}

	for i := 0; i < y; i++ {
		var userId int
		var menuName string
		fmt.Fscan(r, &userId, &menuName)
		user := userList[userId-1]
		switch menuName {
		case "0":
			user.beer()
		case "A":
			user.checkdOut()
			fmt.Fprintln(w, user.getTotalPrice())
		case "food":
			var price int
			fmt.Fscan(r, &price)
			user.food(price)
		case "alcohol":
			var price int
			fmt.Fscan(r, &price)
			user.alcohol(price)
		case "softdrink":
			var price int
			fmt.Fscan(r, &price)
			user.softdrink(price)
		}
	}

	var checkedOutTotalCount int
	for _, v := range userList {
		if v.isCheckedOut {
			checkedOutTotalCount++
		}
	}

	fmt.Fprintln(w, checkedOutTotalCount)
}

func (o *orderUser) food(price int) {
	o.totalPrice += price

	if o.hasOrderedAlcohol {
		o.totalPrice -= 200
	}
}

func (o *orderUser) softdrink(price int) {
	o.totalPrice += price
}

func (o *orderUser) alcohol(price int) {
	if 20 <= o.age {
		o.totalPrice += price
		o.hasOrderedAlcohol = true
	}
}

func (o *orderUser) beer() {
	o.alcohol(500)
}

func (o *orderUser) checkdOut() {
	o.isCheckedOut = true
}

func (o *orderUser) getTotalPrice() int {
	return o.totalPrice
}
