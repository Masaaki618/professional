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

type bank struct {
	pinCode string
	balance int
}

func NewBank(pinCode string, balance int) *bank {
	return &bank{
		pinCode: pinCode,
		balance: balance,
	}
}

func (b bank) validatePinCode(pinCode string) bool {
	return b.pinCode == pinCode
}

func (b *bank) withdraw(amount int) {
	b.balance -= amount
}

func main() {
	defer w.Flush()
	var x, y int
	fmt.Fscan(r, &x, &y)
	bankMap := make(map[string]*bank)
	bankNameList := make([]string, x)
	for i := 0; i < x; i++ {
		var bankName, pinCode string
		var balance int
		fmt.Fscan(r, &bankName, &pinCode, &balance)
		bankMap[bankName] = NewBank(pinCode, balance)
		bankNameList[i] = bankName
	}

	for i := 0; i < y; i++ {
		var bankName, pinCode string
		var balance int
		fmt.Fscan(r, &bankName, &pinCode, &balance)
		account := bankMap[bankName]
		if account.validatePinCode(pinCode) {
			account.withdraw(balance)
		}
	}

	for _, v := range bankNameList {
		fmt.Fprintln(w, v, bankMap[v].balance)
	}
}
