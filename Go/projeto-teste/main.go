package main

import "log"


func main (){
	log.Fatal(novoServidor(1).ListenAndServe())
}


