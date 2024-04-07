/*package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
)

func handlerSims(w http.ResponseWriter, r *http.Request) {
	n1 := Numero{12}
   enviandoNumero, _ := json.Marshal(n1)
   w.Write(enviandoNumero)
}


func multiplica(x int) int{
	return x*2 
}


type Numero struct{
	Num int `json:"Num"`

}
func main() {
	
	
	

   http.HandleFunc("/", handlerSims)
	fmt.Println("colocando servidor no ar!")
    log.Fatal(http.ListenAndServe(":8080", nil))
}*/