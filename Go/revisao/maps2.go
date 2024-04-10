package main

import "fmt"

func main (){

	var mapa = map[string]string{
		"ola" : "show",
		"sair" : "saindo",
	}
	mapa["ue"] = "auau"

	var m string
	
	fmt.Scanln(&m)
	fmt.Println(mapa[m])

}