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
}

func NewOrderUser(age int) *orderUser {
	return &orderUser{
		age: age,
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
		var userId, price int
		var menuName string
		fmt.Fscan(r, &userId, &menuName, &price)
		user := userList[userId-1]
		switch menuName {
		case "food":
			user.food(price)
		case "alcohol":
			user.alcohol(price)
		case "softdrink":
			user.softdrink(price)
		}
	}

	for _, v := range userList {
		fmt.Fprintln(w, v.totalPrice)
	}
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
