package main

import "fmt"

type Desconto interface {
	Aplicar(preco float64) float64
}

type Desconto10 struct{}
type Desconto20 struct{}

func (d Desconto10) Aplicar(preco float64) float64 {
	novoPreco := preco * 0.90
	return novoPreco
}

func (p Desconto20) Aplicar(preco float64) float64 {
	novoPreco := preco * 0.80
	return novoPreco
}

func genericaInterface(nome string, d Desconto, preco float64) {
	resultado := d.Aplicar(preco)
	fmt.Printf("Preço com %s: R$%.2f\n", nome, resultado)
}

func main() {
	valorDesconto10 := Desconto10{}

	valorDesconto20 := Desconto20{}

	genericaInterface("10% de desconto", valorDesconto10, 100.00)
	genericaInterface("20% de desconto", valorDesconto20, 100.00)
}

/*Descrição:
Crie uma interface Desconto com o método:

Aplicar(preco float64) float64

Implemente em structs:

Desconto10 → aplica 10%

Desconto20 → aplica 20%

Teste no main com um preço fixo.*/
