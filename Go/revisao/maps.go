package main

import (
	"errors"
	"fmt"
)

func divide(n int, n2 int) (int, error) {

	if n2 == 0 {
		return 0, errors.New("Não pode dividir por zero")
	}
	return n / n2, nil
}

func soma(n1 int, n2 int) (int, error) {
	var err error
	if err != nil {
		return 0, errors.New("tem que inserir um numero inteiro")
	}
	return n1 + n2, nil
}
func subtracao(n1 int, n2 int) (int, error) {
	var err error
	if err != nil {
		return 0, errors.New("tem que ser um inteiro")
	}
	return n1 - n2, nil
}

func main() {
	var mapa = map[int]func(int, int) (int, error){
		1: divide,
		2: soma,
		3: subtracao,
	}
	var opcao int
	var n1, n2 int
	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&n1)
	fmt.Println("digite o segundo numero")
	fmt.Scanln(&n2)
	fmt.Println("escolha uma operacao: ")
	fmt.Println("1-divisão, 2-soma, 3-sub")
	fmt.Scanln(&opcao)
	fmt.Println(mapa[opcao](n1, n2))

}
