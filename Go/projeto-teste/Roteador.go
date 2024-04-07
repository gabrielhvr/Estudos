package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
func PrimeiroHandle(w http.ResponseWriter,r *http.Request){
	var f1,f2,f3,f4,f5,f6 Filme

	f1 = Filme{Nome: "matriz", Ano: 2002}
	f2 = Filme{Nome: "chuck norris", Ano: 1984}
	f3 = Filme{Nome: "iluminado", Ano: 1983}
	f4 = Filme{Nome: "todo mundo em panico", Ano: 2001}
	f5 = Filme{Nome: "planeta dos macacos", Ano: 2010}
	f6 = Filme{Nome: "america pie", Ano: 2002}

	filmes := []Filme{f1,f2,f3,f4,f5,f6}

	

	disputas := &DisputaFilme{}
	resultadoDisputa, err := disputas.EmparelharFilmesParaDisputas(filmes)
	//TratamentoErro(err)
	if err != nil {
        http.Error(w, "Erro: "+err.Error(), http.StatusInternalServerError)
        return
    }
	b, _ := json.Marshal(&resultadoDisputa)
	w.Write(b)

}
func teste(){
	fmt.Println("testando o servidor")
	http.HandleFunc("/", PrimeiroHandle)




	log.Fatal(http.ListenAndServe(":8080", nil))
}