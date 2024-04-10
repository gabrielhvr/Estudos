package main

import (
	"fmt"
	"strings"
)

func main() {

	var ValidPaths = []string{"/championchip", "/main"}
	var mensagem string
	mensagem = "ola"
	fmt.Println(mensagem)
	tratada := mensagem[1:]
	fmt.Println(tratada)
	fmt.Println(ValidPaths[0])
	fmt.Println(tratamentoCaminho(ValidPaths))

}

func tratamentoCaminho(validPaths []string) []string {
	arrayTratada := []string{}
	for i := 0 ; i < len(validPaths); i++{
		s1 := validPaths[i]
		s2 := s1[1:]
		arrayTratada = append(arrayTratada, s2)
	}
	return arrayTratada
}

func tratamentoArray(texto string) []string {
	return strings.Split(texto, "/")
}
