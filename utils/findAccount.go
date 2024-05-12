package utils

import (
	"log"
	"os"
	"strings"
)

func FoundAcount(account string) bool {
	Found := false

	filePath := "data/users.txt"
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContent := strings.Split(string(file), "\n")

	for i := range fileContent {
		fileDetails := strings.Split(fileContent[i], " ")
		if fileDetails[0] == account {
			return Found
		}
	}
	return Found
}