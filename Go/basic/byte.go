package main

import (
	"fmt"
)

func main() {
	// String original
	str := "Olá, mundo!"

	// Convertendo a string para uma fatia de bytes ([]byte)
	bytes := []byte(str)

	// Imprimindo a fatia de bytes
	fmt.Println("Fatia de bytes:", bytes)

	// Convertendo a fatia de bytes de volta para string
	str2 := string(bytes)

	// Imprimindo a string resultante
	fmt.Println("String resultante:", str2)
}
