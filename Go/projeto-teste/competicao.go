package main

import ("fmt"
		"errors"
)
type DisputaFilme struct{}

func (df DisputaFilme) EmparelharFilmesParaDisputas(filmes []Filme) ([]Disputa,error){
	if len(filmes)%2 != 0{
		return nil, errors.New("a disputa deve ocorre em par")
	}

	disputas := []Disputa{}

	for i := 0; i <len(filmes); i = i +2{
		d:= Disputa{FilmeA: filmes[i], FilmeB : filmes[i + 1]}
		disputas = append(disputas, d)
	}
	return disputas, nil
}


func (c DisputaFilme) ExibirDisputasDeFilmes(disputas []Disputa){
	for i:=0; i < len(disputas); i++{
		d := disputas[i]
		fmt.Printf("%s x %s \n", d.FilmeA.Nome, d.FilmeB.Nome)
	}
}
