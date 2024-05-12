package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"practice/utils"
)

func VerifyLogin(account, pin string) {
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
			userName := fileDetails[1] + " " + fileDetails[2]
			trial, _ := strconv.Atoi(strings.TrimSpace(fileDetails[5]))
			if trial == 0 {
				fmt.Println("Your Pin Is Blocked Please Visit One Of Our Branches To Unlock.")
				os.Exit(0)
			}
			if inputPin == pin {
				fmt.Printf("\nWelcome Back %s!\n", userName)
				HomePage(account, pin, userName)
			} else {

				fmt.Printf("Wrong Pin")
				utils.ResetTrials(account)
				os.Exit(0)
			}
		}
	}
	fmt.Println("Account not found")

	
}