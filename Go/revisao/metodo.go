package main 

import "fmt"

type Pessoa struct{
	Idade int
	Nome string
}

func (p Pessoa) info()int{
	fmt.Println(p.Idade)
	fmt.Println(p.Nome)
	
}

// declaracao do metodo // func (objeto Tipo) nomeDoMetodo() retorno*
// *opcional

func main (){

	var p1 Pessoa
	p1 = Pessoa{22, "gabe"}
	fmt.Println(p1.info())
}