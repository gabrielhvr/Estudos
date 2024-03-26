package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero float64
	var numero2 float64
	var err error
	var divisao float64

	fmt.Print("digite o primeiro numero: ")
	fmt.Scan(&numero)
	fmt.Print("digite o segundo numero: ")
	fmt.Scan(&numero2)
	if numero2 == 0 {
		err = errors.New("Não pode dividir por zero")
	}
	if err != nil{
		fmt.Println("Erro: ", err)
	}
	divisao = numero / numero2
	fmt.Println(divisao)
	fmt.Println(Numero10(11))

}

func Numero10(x int) (int, error){
	if x > 10{
		return 0, errors.New("voce digitou numero maior que 10")
	}
	return x, nil

}

