ackage main

import "fmt"

func main() {
var idade int
var altura float32
var nome string
var tipo bool

nome = "jão"
altura = 186.5
idade = 40
tipo = false
comida := "arroz"
ano := 2024



fmt.Println("Ele é o", nome)
fmt.Println("Ele tem", altura, "altura")
fmt.Println("Ele possui", idade, "ano de idade")
fmt.Println(comida)
fmt.Println("ano: ", ano)


if tipo == true {
	fmt.Println("seu tipo sanguineo é positivo")
}else{
	fmt.Println("seu tipo sanguineo é negativo")
}




}