package main

import "fmt"
type acoesAnimais interface{
	movimento()
	dobra()
}

type Cachorro struct{
	Nome string
}

type Peixe struct{
	Nome string

}

func (c Cachorro) movimento() {
	fmt.Println("O ", c.Nome, "Esta correndo!")
}
func (p Peixe) movimento() {
	fmt.Println("O ", p.Nome, "Esta nadando!")
}


func executarMovimento(a acoesAnimais) {
	a.movimento()
}
func main () {
		c1 := Cachorro{"Holy"}
		p1 := Peixe{"Jubileu"}

		executarMovimento(c1)
		executarMovimento(p1)


}


