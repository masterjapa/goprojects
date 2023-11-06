package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/masterjapa/goprojects/mail-validator/packages/validator"
)

func main() {
	http.HandleFunc("/validate", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := validator.OpenEmailFile("./mail-list/emails.csv")
	emails, _ := validator.ReadEmailFile(file)
	json.NewEncoder(w).Encode(emails)
}
