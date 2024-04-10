package main

import "fmt"

// Definição da interface Colecao
type Colecao interface {
	Adicionar(item string)
	Remover(item string)
	Tamanho() int
}

// Definição do tipo Lista
type Lista struct {
	items []string
}

// Implementação dos métodos da interface Colecao para Lista
func (l *Lista) Adicionar(item string) {
	l.items = append(l.items, item)
}

func (l *Lista) Remover(item string) {
	for i, v := range l.items {
		if v == item {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return
		}
	}
}

func (l *Lista) Tamanho() int {
	return len(l.items)
}

// Definição do tipo Mapa
type Mapa struct {
	items map[string]int
}

// Implementação dos métodos da interface Colecao para Mapa
func (m *Mapa) Adicionar(chave string) {
	m.items[chave] = len(m.items)
}

func (m *Mapa) Remover(chave string) {
	delete(m.items, chave)
}

func (m *Mapa) Tamanho() int {
	return len(m.items)
}

// Função genérica que aceita qualquer tipo que implemente a interface Colecao
func ProcessarColecao(c Colecao) {
	fmt.Println("Tamanho da coleção:", c.Tamanho())

	c.Adicionar("item1")
	c.Adicionar("item2")
	c.Adicionar("item3")

	fmt.Println("Tamanho da coleção após adicionar itens:", c.Tamanho())

	c.Remover("item2")

	fmt.Println("Tamanho da coleção após remover item2:", c.Tamanho())
}

func main() {
	// Criando uma instância de Lista e chamando a função ProcessarColecao
	fmt.Println("Processando Lista:")
	lista := &Lista{}
	ProcessarColecao(lista)

	// Criando uma instância de Mapa e chamando a função ProcessarColecao
	fmt.Println("\nProcessando Mapa:")
	mapa := &Mapa{items: make(map[string]int)}
	ProcessarColecao(mapa)
}
