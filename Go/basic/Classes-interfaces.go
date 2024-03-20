package main 

import "fmt"


func main() {

	var nome string
	var idade int

	pessoa1 := Pessoa{
		Nome: "joão",
		Idade: 30,
	}
	pessoa2 := Pessoa{
		Nome: nome,
		Idade: idade,
	}

	pessoa1.Apresentar()

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)



	fmt.Println()
	fmt.Print(pessoa2.Nome, "  ", pessoa2.Idade)

	pessoa3 := Pessoa{
		Nome : "jose",
		Idade : 13,
	}

	pessoa3.andar()
	pessoa3.ler()
	fmt.Println(pessoa3.computar(3))



}

type Pessoa struct {
	Nome string
	Idade int
}

type action interface{
	andar()
	ler()
	computar()
}

func (p Pessoa) andar(){
	fmt.Println("Estou andando")
}

func (p Pessoa)ler(){
	fmt.Println("Estou lendo")
}

func (p Pessoa)computar(n int)int{
	return n + 10
}




func (p Pessoa) Apresentar() {
	fmt.Printf("olá, eu sou %s e tenho %d anos. \n", p.Nome, p.Idade)
}