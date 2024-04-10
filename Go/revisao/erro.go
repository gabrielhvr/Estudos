package main

import ("fmt"
		"errors"
		"strconv")

func main () {
	/*
	var n1 int
	fmt.Println("digite um numero")
	n1, err := fmt.Scanln(&n1)
	if err != nil {
		errors.New("Digite valor diferente do inteiro")
		fmt.Println("Erro:", err)
	}*/
	var n2 int
	fmt.Println("digite um numero para ver a sua probabilidade")
	_, err := fmt.Scanln(&n2)
	if (err != nil){
		fmt.Println("tem que ser um valor inteiro")
	}

	fmt.Println(probabilidade(n2))
	


}
func probabilidade(p int)(string,error){
	if (p > 100 || p < 0){
		return "Invalido!", errors.New("não existe porcentagem negativa ou maior que 100")
	}
	prob := strconv.Itoa(p)
	return prob + "%", nil
	
}
func divide10(n int)(int, error){
	if (n == 0){
		return 0, fmt.Errorf("Cannot divide by 0")
	}
	return 10/n, nil

}