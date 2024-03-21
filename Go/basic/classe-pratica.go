package main

import "fmt"

func main() {

	empresa1 := Empresa{
		Nome:     "Apple",
		Endereco: "California",
	}

	fmt.Println("Sistema de vendas")
	fmt.Print("Nome da Empresa: ", empresa1.Nome, "\t")
	fmt.Println("Endereço da Empresa: ", empresa1.Endereco)

	fmt.Println("usando as funções")
	fmt.Println(empresa1.getNome())
	empresa1.setNome("Fejoada")
	empresa1.Nome = "Ronald"
	fmt.Println("Nome da Empresa: ", empresa1.Nome, "\t")
	fmt.Println(empresa1.getNome())

}

type Empresa struct {
	Nome     string
	Endereco string
	//Valor    float
	//CNPJ     int
}

func (e *Empresa) setNome(nome string) {
	e.Nome = nome
}

func (e *Empresa) getNome() string {
	return e.Nome
}

func (e *Empresa) setEndereco(endereco string) {
	e.Endereco = endereco
}

func (e *Empresa) getEndereco() string {
	return e.Endereco
}
