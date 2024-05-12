package utils

import (
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ResetTrials(account string) {
	filepath := "data/users.txt"

	// Open the file for reading and writing
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Seek to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Error seeking to the beginning of the file: %v", err)
	}

	// Create a buffer to store the modified content
	buffer := make([]byte, 1024)

	// Read content from the file into the buffer
	n, err := file.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from file: %v", err)
	}

	// Modify the content in the buffer
	content := string(buffer[:n])
	lines := strings.Split(content, "\n")

	// Loop through each line and decrement trials for the specified account if it's not already at 0
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == account {
			// Parse trials as an integer
			trials, err := strconv.Atoi(fields[len(fields)-1])
			if err != nil {
				log.Fatalf("Error parsing trials for account %s: %v", account, err)
			}
			// Decrement trials if it's not already at 0
			if trials > 0 {
				trials--
				fields[len(fields)-1] = strconv.Itoa(trials)
				lines[i] = strings.Join(fields, " ") // Reconstruct the line with decremented trials
			}
			break // No need to continue searching
		}
	}

	// Join the modified lines back together
	modifiedContent := strings.Join(lines, "\n")

	// Truncate the file before writing to it
	err = file.Truncate(0)
	if err != nil {
		log.Fatalf("Error truncating file: %v", err)
	}

	// Seek back to the beginning of the file before writing
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Error seeking to the beginning of the file: %v", err)
	}

	// Write the modified content back to the file
	_, err = file.WriteString(modifiedContent)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Sync the file to ensure changes are flushed to disk
	err = file.Sync()
	if err != nil {
		log.Fatalf("Error syncing file: %v", err)
	}

	// fmt.Println("Trials decremented successfully.")
}
