package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

//var caminhoValidos = []string{"/","/about"}

func handlerJogo(w http.ResponseWriter,r *http.Request){
	J1 := Jogo{
		"cod",2012,
	}
	enviandoJogo, _ := json.Marshal(J1)
	w.Write(enviandoJogo)
}


func handlerAbout(w http.ResponseWriter, r *http.Request){
	fmt.Println("mensagem no terminal")
	fmt.Fprintln(w, "Bem vindo ao About!")
}