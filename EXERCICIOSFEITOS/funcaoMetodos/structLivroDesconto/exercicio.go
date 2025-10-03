package main

import "fmt"

type livro struct {
	titulo string
	preco  float64
}

func (l livro) AplicarDesconto(p float64) float64 {

	Novopreco := l.preco * (1 - p/100)
	return Novopreco
}

func main() {
	livroFantasia := livro{
		titulo: "Harry potter",
		preco:  39.99,
	}

	fmt.Printf("Preço com desconto: R$ %.2f\n", livroFantasia.AplicarDesconto(10.00))
}

/*Crie uma struct Livro com titulo e preco.
Implemente um método AplicarDesconto(p float64) que diminui o preço do livro de acordo com a porcentagem recebida.*/
