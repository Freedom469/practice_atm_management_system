package service

import (
	"fmt"
	"os"
	"practice/manage"
	"practice/utils"
	"strconv"
)


func TransferFunds(accountToDebit string, balance *int) {
	var transferAmount string

	fmt.Print("Enter Account Number: ")
	var accountToCredit string
	fmt.Scan(&accountToCredit)
	fmt.Print("Enter Amount: ")
	fmt.Scan(&transferAmount)

	intAmount, _ := strconv.Atoi(transferAmount)
	if *balance >= intAmount {
		if utils.FoundAcount(accountToCredit){
			manage.CreditAccount(accountToCredit, accountToDebit, transferAmount)
			manage.DebitAccount(accountToDebit, transferAmount)
		} else {
			fmt.Println("Acoount Not Found")
			os.Exit(0)
		}
	} else {
		fmt.Printf("Sorry You Do Not Have Enough Cash In your Account To Transfer %s To Account %s.\n", transferAmount, accountToCredit)
		fmt.Printf("Your Available Balance is %v Please Top Up Your Account and Try Again.", balance)
		os.Exit(0)
	}

}
