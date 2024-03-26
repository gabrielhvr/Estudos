package main

import (
	"fmt"
)


func main() {


	var num int
	result, err := soma(1, 21)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result, err)
	println("hello, world")
	//http.ListenAndServe(":8080", nil)
	nome := "gabriel"
	nome = "sims"
	fmt.Println(nome)
	fmt.Println(num)
	fmt.Println(soma(1,2))
}

func soma(a int, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("soma maior que 10")
	}
	return a + b, nil
}