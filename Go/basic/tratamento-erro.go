package main 
// para tratamento de erro, primeiro importe a biblioteca erro
import ("fmt"
		"errors"
)
func main () {
	a, b := 10, 20

	resultado, err := divide(10, 0)


	fmt.Println(a + b) // exemplo de declaração de multiplas variaveis em golang

	if err != nil{

		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Resultado da divisão:", resultado)



}
// no golang é possivel retorna mais de valor, o segundo retorno geralmente é usado para retorna um error
func divide (x, y int)(int, error) {
	if y == 0 {

		return 0, errors.New("impossível dividir por zero") // retorna um inteiro 0 e notifica o erro
	}
	return x / y, nil // sempre tem que retorna algo, neste caso a divisão de x / y e nil para erro

}

