package main

import (
	"fmt"
)

// criando os atributos
type Pessoa struct {
	Nome   string
	Idade  int
	Altura float32
	Luta   bool
}

// criando os metodos
func (p Pessoa) ImprimirInfo() {
	fmt.Println("Nome: ", p.Nome)
	fmt.Println("Idade: ", p.Idade)
	fmt.Println("Altura: ", p.Altura)
	fmt.Println("Luta: ", p.Luta)
}

func (p *Pessoa) aumentarIdade() {
	p.Idade++
}

func main() {

	pessoa1 := Pessoa{"gabriel", 28, 1.85, true}
	//pessoa2 := Pessoa{"maria", 25, 1.70, false}
	pessoa1.ImprimirInfo()
	pessoa1.aumentarIdade()
	pessoa1.ImprimirInfo()


}
