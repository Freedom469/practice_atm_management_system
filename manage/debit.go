package manage

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func DebitAccount(accountToDebit, amount string) {
	filepath := "data/users.txt"
	var newBalance int

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the content of the file
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from the file: %v", err)
	}
	content := string(buffer[:n])
	lines := strings.Split(content, "\n")

	// Find the account and update the balance
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) >= 5 && parts[0] == accountToDebit {
			balance, err := strconv.Atoi(parts[4])
			if err != nil {
				log.Fatalf("Error converting balance to integer: %v", err)
			}
			amountInt, err := strconv.Atoi(amount)
			if err != nil {
				log.Fatalf("Error converting amount to integer: %v", err)
			}
			newBalance = balance - amountInt
			parts[4] = strconv.Itoa(newBalance)
			lines[i] = strings.Join(parts, " ")
			// pin := parts[3]
			// userName := parts[1] + " " + parts[2]
			// service.HomePage(accountToDebit, pin, userName)
			break
		}
	}

	// Write the updated lines back to the file
	updatedContent := strings.Join(lines, "\n")
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Error seeking the beginning of the file: %v", err)
	}
	_, err = file.WriteString(updatedContent)
	if err != nil {
		log.Fatalf("Error writing to the file: %v", err)
	}

	fmt.Printf("You have successfully transferred %s to account %s. Balance: %d\n", amount, accountToDebit, newBalance)
}