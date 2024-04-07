package main

import ("net/http"
		"encoding/json"
		"time"
		"fmt"
	
	)

var diretoriosValidos = []string{"/main","/ola"}


//criar uma verificação de diretorios
func verficadiretorio (diretorio string)bool{
	for i:= 0; i <len(diretoriosValidos); i++{
	if (diretorio == diretoriosValidos[i]){
		return true
	}
	}
	return false
	
}

type Router int



func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !verficadiretorio(r.URL.Path){
		w.WriteHeader(418)
		w.Write([]byte("{\"codigo\": 418, \"mensagem\": caminho invalido}"))
	}

	
	if r.URL.Path == "/main"{
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
	if r.URL.Path == "/ola"{
		fmt.Fprint(w, "fudeu")
	}

}



func novoServidor(r Router) *http.Server {
	return &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        r,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}
}