package main

type disputador interface {

	EmparelharFilmesParaDisputas(filmes []Filme) []Disputa
}
