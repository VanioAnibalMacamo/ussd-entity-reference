package ussd

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

var (
	phoneNumber string
	valor       string
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		sessionId := r.FormValue("sessionId")
		serviceCode := r.FormValue("serviceCode")
		text := r.FormValue("text")

		fmt.Println("Session ID:", sessionId)
		fmt.Println("Service Code:", serviceCode)

		var response string

		switch text {
		case "":
			response = "Mkesh \n"
			response += "1. Requisitar Referencia \n"
		case "1":
			response = "Digite o número do Telefone\n"
		default:
			if strings.HasPrefix(text, "1*") {
				// Verificar se o texto começa com "1*" (indicando que o usuário digitou o número do telefone)
				phoneNumber = strings.TrimPrefix(text, "1*")
				if len(phoneNumber) != 9 {
					response = "END O número de telefone deve ter 9 dígitos\n"
				} else {
					response = "END Seu número de telefone é " + phoneNumber + ". Digite o valor\n"
				}
			} else if phoneNumber != "" {
				// Se phoneNumber estiver definido, significa que o usuário já digitou o número do telefone
				valor = text
				response = "END Seu número de telefone é " + phoneNumber + " e o valor digitado é " + valor
			} else {
				response = "END Seleção inválida"
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, response)
	} else {
		http.Error(w, "Apenas requisições POST são suportadas", http.StatusMethodNotAllowed)
	}
}
