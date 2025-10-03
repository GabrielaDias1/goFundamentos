package main

import "fmt"

// cria struct de produto, produto com caracteristicas que fazem sentido estarem juntos
// Struct Produto com campos relacionados (nome, preço e estoque)
type Produto struct {
	Nome      string
	Preco     float64
	EmEstoque bool
}

func main() {

	estoqueProduto := []Produto{
		//// Cria uma slice de produtos
		// assim crio varios produtos para o slice poder percorrer
		//Criar uma coleção de objetos assim dessa forma por conta do slice para armazenar vários e usar for range para percorrer.
		{Nome: "Candida", Preco: 8.19, EmEstoque: true},
		{Nome: "Veja", Preco: 8.69, EmEstoque: false},
		{Nome: "Desifetante", Preco: 5.99, EmEstoque: true},
	}
	for _, p := range estoqueProduto {
		fmt.Printf("Produto: %s, Preço: %.2f, Em estoque: %t\n", p.Nome, p.Preco, p.EmEstoque)
	}
}

/*Lista de produtos

Crie uma slice de Produto e adicione 3 produtos.
Use um for para imprimir todos os produtos no console.*/
