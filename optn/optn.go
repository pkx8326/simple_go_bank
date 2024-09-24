package optn

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balancefile = "Balance.txt"

func Menu() (c string) {
	fmt.Println("Please choose one of the following transactions :")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
	fmt.Print("Your choice :")
	fmt.Scan(&c)
	return
}

func Menu2() (c string) {
	fmt.Println("Please choose one of the following transactions :")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Exit")
	fmt.Print("Your choice :")
	fmt.Scan(&c)
	return
}

func Readbalance() (balance float64, FileNotFoundErr error) {
	balanceData, err := os.ReadFile(balancefile)
	if err != nil {
		fmt.Println("No balance file found. Please consider making a deposit with out bank.")
		FileNotFoundErr = errors.New("filenotfound")
		return
	} else {
		balanceText := string(balanceData)
		balance, _ = strconv.ParseFloat(balanceText, 64)
		return
	}
}

func Writebalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balancefile, []byte(balanceText), 0644)
}

func Checkbalance(balance float64) {
	fmt.Println("Your current balance is", balance)
}

func Deposit(balance float64) {
	var amount float64
	for {
		fmt.Print("Please input an amount to deposit :")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid input.")
			continue
		} else {
			updatedBalance := balance + amount
			fmt.Println("Your current balance is", updatedBalance)
			Writebalance(updatedBalance)
			break
		}

	}
}

func Withdraw(balance float64) {
	var amount float64
	for {
		fmt.Print("Please input an amount to withdraw :")
		fmt.Scan(&amount)

		if amount <= 0 {
			fmt.Println("Invalid input.")
		} else {
			if amount > balance {
				fmt.Println("Insufficient funds.")
				break

			} else {
				updatedBalance := balance - amount
				fmt.Println("Your current balance is", updatedBalance)
				Writebalance(updatedBalance)
				break
			}
		}

	}
}

func Nexttransaction() (nxttxn string) {
	for {
		fmt.Print("Do you have another transaction? [Y/N] :")
		fmt.Scan(&nxttxn)
		if nxttxn == "Y" || nxttxn == "y" || nxttxn == "N" || nxttxn == "n" {
			return
		} else {
			fmt.Println("Invalid input.")
			continue
		}
	}
}

func Exit() (nexttxn string) {
	fmt.Println("Thank you for choosing our bank. Goodbye.")
	nexttxn = "n"
	return
}
