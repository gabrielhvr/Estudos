package main

import "fmt"

func main() {
	var comando int
	var consulta = map[int]string{
		1: "VOCÊ É CHARMOSO!",
		2: "VOCÊ É BONITO!",
	}

	//var mensagem string

	fmt.Println("Digite um comando: ")
	fmt.Println("1 - você é charmoso!")
	fmt.Println("2 - você é bonito!")
	fmt.Scan(&comando)

	fmt.Println(consulta[comando])

}
