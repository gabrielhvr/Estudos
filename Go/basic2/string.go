package main

import "fmt"


func removerCaractere(str string, index int) string {
    return str[:index] + str[index+1:]
}



func main () {
	
	var nome string = "ggabriel"
	
	fmt.Println(string(nome[0])) // ver um caracter de uma string
    fmt.Println("String original:", nome)

    // Removendo o primeiro caractere
   // novoNome := removerCaractere(nome, 0)
   novoNome := nome[1:]

    fmt.Println("String modificada:", novoNome)

}
