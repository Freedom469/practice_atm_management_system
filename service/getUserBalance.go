package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"practice/utils"
)

func GetUserBalance(account, pin string) int {
	filePath := "data/users.txt"
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.Split(string(file), "\n")

	for i := range fileContent {
		fileDetails := strings.Split(fileContent[i], " ")
		if fileDetails[0] == account {
			inputPin := fileDetails[3]
			// userName := fileDetails[1] + " " + fileDetails[2]
			balance, _ := strconv.Atoi(fileDetails[4])
			trial, _ := strconv.Atoi(strings.TrimSpace(fileDetails[5]))
			if trial == 0 {
				fmt.Println("Your Pin Is Blocked Please Visit One Of Our Branches To Unlock.")
				os.Exit(0)
			}
			if inputPin == pin {
				// fmt.Printf("Welcome %s!\n", userName)
				// fmt.Printf("\nYour Available Balance: %s Ksh.\n", balance)
				return balance
			}
			fmt.Printf("Wrong Pin Tries left %d.\n", trial -1 )
			utils.ResetTrials(account)
			 // Return empty user if PIN is wrong
		}
	}

	fmt.Println("Account not found")
	return 0 // Return empty user if account is not found
}

