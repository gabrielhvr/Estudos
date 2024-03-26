package main

import "fmt"

type Livro struct {
	Titulo   string
	Autor    string
	N_Pagina int
}

func (l *Livro) setLivro(titulo string, autor string, n_pagina int) {
	l.Titulo = titulo
	l.Autor = autor
	l.N_Pagina = n_pagina
}

func (l *Livro) getTitulo() string {
	return l.Titulo
}

func (l *Livro) getAutor() string {
	return l.Autor
}

func (l *Livro) getPagina() int {
	return l.N_Pagina
}

func imprimirLivro(l Livro){
	fmt.Println("Titulo do livro:",l.getTitulo())
	fmt.Println("Nome do Autor:",l.getAutor())
	fmt.Println("Numero de Paginas:", l.getPagina())
}