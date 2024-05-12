package manage

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreditAccount(accountTo, accountFrom, amount string) error {
	filePath := "data/users.txt"
	var newBalance int

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the content of the file
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	fileContent := string(buffer[:n])
	lines := strings.Split(fileContent, "\n")

	found := false
	// Find the account and update the balance
	for i, line := range lines {
		parts := strings.Fields(line) // Split by any whitespace
		if len(parts) >= 5 && parts[0] == accountTo {
			balance, err := strconv.Atoi(parts[4])
			if err != nil {
				return fmt.Errorf("error converting balance to integer: %v", err)
			}
			amountInt, err := strconv.Atoi(amount)
			if err != nil {
				return fmt.Errorf("error converting amount to integer: %v", err)
			}
			newBalance = balance + amountInt
			parts[4] = strconv.Itoa(newBalance)
			lines[i] = strings.Join(parts, " ")
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("account not found: %s", accountTo)
	}

	// Write the updated lines back to the file
	updatedContent := strings.Join(lines, "\n")
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking the beginning of the file: %v", err)
	}
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating file: %v", err)
	}
	if _, err := file.WriteString(updatedContent); err != nil {
		return fmt.Errorf("error writing to the file: %v", err)
	}

	fmt.Printf("Confirm You have Received %s From %s. Your New Balance: %d\n", amount, accountFrom, newBalance)
	return nil
}
