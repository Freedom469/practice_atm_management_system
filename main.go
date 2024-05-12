package main

import (
	"fmt"
	"log"
	"practice/service"
	"practice/utils"
)

func main() {
	fmt.Println("##################################################################")
	fmt.Println("\n\t\t Please Sign Into Your Account to Continue")
	fmt.Println()
	fmt.Print("Enter Your Account Number: ")
	var account string
	less := false
	fmt.Scan(&account)
	if len(account) != 8 {
		less = true
		for less {
			utils.ClearScreen()
			fmt.Print("Please Enter a Valid account Number (length-8): ")
			_, err := fmt.Scan(&account)
			if err != nil {
				log.Fatal(err)
			}
			if len(account) == 8 {
				less = false
			}
		}
	}
	var pin string
	if len(account) == 8 {
		fmt.Print("\nEnter Your PIN: ")
		fmt.Scan(&pin)

		if len(pin) != 4 {
			less = true
			for less {
				utils.ClearScreen()
				fmt.Print("Please Enter a Valid PIN Number (length-4): ")
				_, err := fmt.Scan(&pin)
				if err != nil {
					log.Fatal(err)
				}
				if len(pin) == 4 {
					less = false
				}
			}
		}

	}
	fmt.Println("##################################################################")
	utils.ClearScreen()
	service.VerifyLogin(account, pin)
}
