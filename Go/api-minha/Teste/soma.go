package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
    // Extrair os parâmetros da URL
    query := r.URL.Query().Get("numbers")

    // Separar os números
    query = strings.Replace(query, "+", " ", -1)
	numbers := strings.Fields(query)

    // Inicializar a variável para armazenar o resultado da soma
    sum := 0

    // Calcular a soma dos números
    for _, numStr := range numbers {
        num, err := strconv.Atoi(numStr)
        if err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }
        sum += num
    }

    // Escrever a resposta como texto simples
    fmt.Fprintf(w, "A soma dos números é: %d", sum)
}

func main() {
    // Mapear a rota para a função de manipulador
    http.HandleFunc("/sum", sumHandler)

    // Iniciar o servidor na porta 8080
    fmt.Println("Servidor iniciado na porta :8080")
    http.ListenAndServe(":8080", nil)
}
