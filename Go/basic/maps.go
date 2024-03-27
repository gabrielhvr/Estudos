package main

import "fmt"

func main() {
	nomes := []string{"tiago", "Dani", "Marcos", "Maria"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Rafael")
	fmt.Println(nomes, len(nomes), cap(nomes))
	maps()



}
func maps(){
	idades := make(map[string]uint8)
	idades ["thiago"] = 31
	idades ["dani"] = 36
	idades ["Maria"] = 23
	
	fmt.Println(idades) // o maps não segue a ordem
}


