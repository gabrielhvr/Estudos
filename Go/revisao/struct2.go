package main
/*
Claro! Aqui está um exercício simples para praticar o conceito de structs compostas:

Imagine que você está desenvolvendo um sistema para uma loja de eletrodomésticos. Crie três structs: Produto, Cliente e Compra. A struct Produto deve conter informações sobre o produto, como nome e preço. A struct Cliente deve conter informações sobre o cliente, como nome e endereço. A struct Compra deve representar uma compra realizada pelo cliente e deve conter um produto comprado e o cliente que fez a compra.

Com base nessas structs, você pode criar algumas instâncias de produtos e clientes e simular uma compra, criando uma instância da struct Compra que inclui o produto comprado e o cliente que realizou a compra.*/

import "fmt"

type Product struct{
	NameProduct string
	Price int
}

type Person struct{
	NamePerson string
	Address string
}

type Purchase struct{
	Product
	Person
}


func main () {
	var p1,p3 Purchase
	p1.NameProduct = "laptop"
	p1.Price = 22
	p1.NamePerson = "luiz"
	p1.Address = "3030 greenwich village"
	p2 := Purchase{
		Product{NameProduct : "cpu", Price : 22,},
		Person{NamePerson : "rafael", Address : "shql 202 casa 12",},
		}
	p3 = Purchase{
	Product{"geladeira",22},
	Person{"Roberta", "lago norte"},}
		

	fmt.Println("Primeira compra realizada!\nNome do cliente: ",p1.NamePerson,"\nProduto : ", p1.NameProduct,"\nPreco do produto: ",p1.Price, "\nEndereço de entrega: ",p1.Address)
	fmt.Println("Segunda compra realizada!\nNome do cliente: ",p2.NamePerson,"\nProduto : ", p2.NameProduct,"\nPreco do produto: ",p2.Price, "\nEndereço de entrega: ",p2.Address)
	
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
