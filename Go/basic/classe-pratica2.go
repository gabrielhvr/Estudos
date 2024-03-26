package main 

import "fmt"
type Pessoa2 struct{
	Nome string
	Idade int
}


func main (){
	var nome string
	var idade int
	var alguem Pessoa2
	//var varios []Pessoa2
	fmt.Println("nome:")
	fmt.Scan(&nome)
	fmt.Println("idade:")
	fmt.Scan(&idade)
	alguem = Pessoa2{Nome: nome, Idade : idade}
	fmt.Println("Nome:", alguem.Nome)
	fmt.Println("Idade:", alguem.Idade)



}

