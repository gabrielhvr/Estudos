package main

import "fmt"

func main() {

	// imprimi o conteudo dentro de uma slice
	nomes :=[]string{"joao", "bob", "carol", "julia"}
	for _, nome := range nomes { // _ aqui onde colocar interador com i por exemplo mas neste caso colocamos _ para não precisar usar essa variavel
		fmt.Println(nome)
	}
	for i, elemento := range nomes { // com i
		fmt.Println("nome:",elemento, "index:", i)
}}