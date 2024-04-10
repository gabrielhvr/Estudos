package main

import "fmt"

//criar uma estrutura que representa um carro, que inclui informações sobre o modelo do carro e o motorista:
type CarroModelo struct {
	Ano int
	NomeModelo string

	
}

type Motorista struct{
	Nome string
	Idade int
}

type Carro struct {
	CarroModelo
	Motorista
}


func main () {
	var c1 Carro
	c1.Ano = 2012
	c1.NomeModelo = "civic"
	c1.Nome = "Gabriel"
	c1.Idade = 28
	fmt.Println(c1.Ano, c1.NomeModelo, c1.Nome, c1.Idade)


}
