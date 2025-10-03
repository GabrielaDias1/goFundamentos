package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

type Pedido struct {
	Numero  int
	Produto []Produto
}

func main() {

	pedido := Pedido{
		Numero: 123321,
		Produto: []Produto{
			{Nome: "Hambúrguer", Preco: 19.99},
			{Nome: "Batata", Preco: 11.00},
		},
	}
	fmt.Printf("Pedido número %d\n", pedido.Numero)
	fmt.Println("Produtos solicitados:")
	for _, p := range pedido.Produto {
		fmt.Printf("- %s: R$ %.2f\n", p.Nome, p.Preco)
	}
}

/*Pedido com produtos

Crie uma struct Produto com Nome e Preco.

Crie uma struct Pedido com Numero (int) e uma slice de Produto.

Crie um pedido com 2 produtos e imprima o número do pedido e os produtos com seus preços.*/
