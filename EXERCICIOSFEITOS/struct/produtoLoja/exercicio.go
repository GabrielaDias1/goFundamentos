package main

import "fmt"

// crio uma struct, agrupamento de dados que faz sentido juntos, em estoque faz sentido como oque tem um produto
type Produto struct {
	Nome      string
	Preco     float64
	EmEstoque bool
}

func main() {
	// aqui crio uma variavel que vai ser do tipo produto, assim criando um produto inserindo suas informações
	//Cria uma variável 'produto' do tipo Produto e inicializa com nome, preço e estoque
	produto := Produto{
		Nome:      "veja",
		Preco:     8.69,
		EmEstoque: true,
	}

	fmt.Printf("O produto %s custa %.2f e está em estoque = %t\n", produto.Nome, produto.Preco, produto.EmEstoque)

}

/*Produto de loja

Crie uma struct Produto com os campos Nome (string), Preco (float64) e EmEstoque (bool).
Crie um produto e exiba a mensagem:
"O produto X custa Y e está em estoque = Z".*/
