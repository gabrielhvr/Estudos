package main

import (
    "fmt"
)

func main() {
	var condicao bool
	condicao = true
	
	var opcao int = 1

	for condicao == true {
		fmt.Print("Digite um valor: ")
        fmt.Scan(&opcao)
        if opcao == 0 {
            fmt.Println("fechou o programa")
			condicao = false
        }else{
            fmt.Println("ta o programa")
        }
	}


}
