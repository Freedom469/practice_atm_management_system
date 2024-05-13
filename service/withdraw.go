package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"practice/utils"
)

func WithdrawFunds(acoount string) {
	filePath := "data/users.txt"
	var withdrawAmount int

	file, fileOpenError := os.OpenFile(filePath, os.O_RDWR, 0644)

	if fileOpenError != nil {
		log.Fatalf("Error Opening File: %v", fileOpenError)
	}

	buffer := make([]byte, 1024)

	n, fileReadError := file.Read(buffer)

	if fileReadError != nil {
		log.Fatalf("Error While Reading The file %v", fileReadError)
	}

	fileContent := string(buffer[:n])

	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		for i, part := range parts {
			if i == 0 && part == acoount {
				fmt.Print("Enter Amount: ")
				fmt.Scan(&withdrawAmount)
				balance, _ := strconv.Atoi(parts[4])
				if withdrawAmount >= balance {
					fmt.Printf("Sorry You Do Not Have Enough Balance To Withdraw %d\n", withdrawAmount)
					fmt.Printf("Your Available Balance is %d\n", balance)
					break
				}
				fmt.Printf("Confirm Withdrawal of %d (y/n): ", withdrawAmount)
				var Confirm string
				fmt.Scan(&Confirm)
				if Confirm == "y" || Confirm == "Y" {
					utils.ClearScreen()
					fmt.Printf("Congratulations Withrawal Of Ksh %d Was Successful!\n", withdrawAmount)
					fmt.Println("Facilitation Fee : 30")
					fmt.Printf("Available Balance: %d.\n", balance-withdrawAmount-30)
					break
				} else {
					continue
				}
				break
			}
		}
	}
}