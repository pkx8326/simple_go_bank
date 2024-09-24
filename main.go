package main

import (
	"fmt"

	"example.com/gobank/optn"
)

func main() {
	var nexttxn, c string
	fmt.Println("Welcome to GO Bank.")
	for {
		balance, err := optn.Readbalance()
		if err != nil {
			c = optn.Menu2()
			switch c {
			case "1":
				optn.Checkbalance(balance)
				nexttxn = optn.Nexttransaction()
			case "2":
				optn.Deposit(balance)
				nexttxn = optn.Nexttransaction()
			case "3":
				nexttxn = "n"
			}
		} else {
			c = optn.Menu()

			switch c {
			case "1":
				optn.Checkbalance(balance)
				nexttxn = optn.Nexttransaction()
			case "2":
				optn.Deposit(balance)
				nexttxn = optn.Nexttransaction()
			case "3":
				optn.Withdraw(balance)
				nexttxn = optn.Nexttransaction()
			case "4":
				nexttxn = "n"
			}
		}

		if nexttxn == "Y" || nexttxn == "y" {
			continue
		} else {
			optn.Exit()
			break
		}
	}
}

/////////////////////////////////////////////////////////////////////////////
