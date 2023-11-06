package validator

import (
	"encoding/csv"
	"log"
	"os"
)

func OpenEmailFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening emails.csv ", err)
		return nil, err
	}
	return file, nil
}

func ReadEmailFile(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	emails, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading emails.csv ", err)
		return nil, err
	}
	defer file.Close()
	return emails, nil
}
