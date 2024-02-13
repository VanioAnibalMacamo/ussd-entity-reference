package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleRequest)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão se a variável de ambiente não estiver definida
	}
	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		sessionId := r.FormValue("sessionId")
		serviceCode := r.FormValue("serviceCode")
		phoneNumber := r.FormValue("phoneNumber")
		text := r.FormValue("text")

		fmt.Println("Session ID:", sessionId)
		fmt.Println("Service Code:", serviceCode)

		var response string

		switch text {
		case "":
			response = "CON What would you want to check \n"
			response += "1. My Account \n"
			response += "2. My phone number"
		case "1":
			response = "CON Choose account information you want to view \n"
			response += "1. Account number \n"
		case "2":
			response = "END Your phone number is " + phoneNumber
		case "1*1":
			accountNumber := "ACC1001"
			response = "END Your account number is " + accountNumber
		// Add additional cases for handling other combinations of menu selections
		default:
			response = "END Invalid selection"
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, response)
	} else {
		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
	}
}
