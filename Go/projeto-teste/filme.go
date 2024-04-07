package main

import ("fmt")

func main (){
	var f1,f2,f3,f4 Filme

	f1 = Filme{Nome: "matriz", Ano: 2002}
	f2 = Filme{Nome: "chuck norris", Ano: 1984}
	f3 = Filme{Nome: "iluminado", Ano: 1983}
	f4 = Filme{ Nome: "todo mundo em panico", Ano: 2001}

	Filmes := []Filme{f1,f2,f3,f4}

	

	cinema1 := NovoCinema(Filmes)

	fmt.Println(cinema1)
	



}




type Filme struct {
	Nome string
	Ano int
}


func (f Filme) getNome() string{
	return f.Nome
}


type Cinema struct{
	Filmes []Filme
}

func NovoCinema (filmes []Filme) *Cinema{
	c := &Cinema{}
	c.Filmes = filmes
	return c
	
}

