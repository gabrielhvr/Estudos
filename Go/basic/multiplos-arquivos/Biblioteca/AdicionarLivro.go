package main

import "fmt"

func adicionarLivro(livro []Livro) []Livro {

	var livroAdicionado Livro
	var titulo, autor string
	var n_pagina int
	fmt.Println("Digite o livro que deseja adicionar: ")
	fmt.Scan(&titulo)
	fmt.Println("Digite o autor: ")
	fmt.Scan(&autor)
	fmt.Println("Digite o numero de paginas: ")
	fmt.Scan(&n_pagina)

	livroAdicionado.setLivro("harry", "jose", 250)
	livro = append(livro, livroAdicionado)
	return livro
}
