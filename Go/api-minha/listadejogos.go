package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {

    http.HandleFunc("/", handlerJogo)
    http.HandleFunc("/about", handlerAbout)
    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Jogo struct {
    Nome string `json:"Nome"`
    Ano  int    `json:"Ano"`
}


