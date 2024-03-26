package main

import (
	"fmt"
)
//var Biblioteca []Livro

func main() {
	var harry Livro
	var Biblioteca []Livro

	harry.setLivro("harry", "jose", 250)

	fmt.Println(harry.getTitulo())
	//adicionarLivro(Biblioteca)
	//adicionarLivro(Biblioteca)

	//imprimirLivro(harry)
	Biblioteca = adicionarLivro(Biblioteca)
	fmt.Println(Biblioteca)	
}

