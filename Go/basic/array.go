package main

import (
	"fmt"
)

func main() {
	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50
	
	fmt.Println("Tamanho do array: ", len(numeros))
	imprimirArray(numeros)
	fmt.Println()
	fmt.Println("O maior valor é :",maiorValor(numeros))

}



func imprimirArray(numero [5]int){
	for i := 0; i < len(numero); i++{
		fmt.Print(numero[i], " ")
	}
	fmt.Print("\t")
}

func maiorValor(numero [5]int)int{
	var x int
	x = numero[0]
	for i:= 0; i < len(numero) - 1; i++{
		if numero[i] < numero[i + 1]{
			x = numero[i + 1]
		}
	}
	return x


}