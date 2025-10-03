package main

import "fmt"

type produto struct {
	nome  string
	preco float64
}

//criacao struct

// criando parametro percentual
func percentualDesconto(p produto, percentual float64) float64 {
	desconto := p.preco * (percentual / 100)
	precoComDesconto := p.preco - desconto
	return precoComDesconto
}

func main() {
	produtoVeja := produto{
		nome:  "Veja",
		preco: 9.99,
	}

	precoFinal := percentualDesconto(produtoVeja, 10)
	fmt.Println("Produto com desconto:", precoFinal)

}

/*Cálculo de desconto em produto

Crie uma struct Produto com Nome e Preco.
Crie uma função que receba um produto e um percentual de desconto (float64) e retorne o preço com desconto.
No main, crie um produto, aplique 10% de desconto e imprima o resultado.*/
