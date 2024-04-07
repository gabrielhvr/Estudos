package main

import "fmt"

func main (){


	x := 12

	func (x int){ // a funcao não tem nome

		fmt.Println(x*2)
	}(x) // (x) equivale a chamar a funcao



}