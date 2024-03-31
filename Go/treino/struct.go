package main

import (
	"errors"
	"fmt"
)

type Cachorro struct {
	idade int
	nome  string
}

func divisao(n, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("Não pode dividir por 0")
	}
	return n / n2, nil

}

func main() {
	var husky Cachorro
	husky.idade = 12
	husky.nome = "holy"

	fmt.Println(husky.idade, "---", husky.nome)
	resultado, err := divisao(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)

}

//fmt.Println("Resultado da divisão:", resultado)
