package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func main() {
	router := Router{}

	// Mapeamento de URL para funções
	rotas := map[string]func(http.ResponseWriter, *http.Request){
		"/main": router.handleMain,
		"/ola":  router.handleOla,
	}

	http.Handle("/", router)

	fmt.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", nil)
}

func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Verifica se a URL está mapeada no mapa de rotas
	if funcao, ok := rotas[r.URL.Path]; ok {
		// Se estiver, chama a função correspondente
		funcao(w, r)
	} else {
		// Se não, retorna um erro 404
		http.NotFound(w, r)
	}
}

func (Router) handleMain(w http.ResponseWriter, r *http.Request) {
	// Sua lógica para a rota "/main" aqui
	fmt.Fprint(w, "Lógica para a rota /main")
}

func (Router) handleOla(w http.ResponseWriter, r *http.Request) {
	// Sua lógica para a rota "/ola" aqui
	fmt.Fprint(w, "Olá, mundo!")
}
