package main

import "fmt"

type Carro struct {
	Nome string
	Cor  string
	Ano  int
}

func (c Carro) andar() {
	fmt.Printf("O carro %s da cor %s do ano %d está andando.\n", c.Nome, c.Cor, c.Ano)
}

func (a Carro) buzinar() {
	fmt.Printf("Bibii!\n")
}

func (b Carro) freiar() {
	fmt.Printf("O carro %s está freiando!\n", b.Nome)
}

func main() {
	hondaCivic := Carro{
		Nome: "Honda Civic",
		Cor:  "prata",
		Ano:  2020,
	}
	hondaCivic.andar()
	hondaCivic.buzinar()
	hondaCivic.freiar()
}

/* carro
struct Carro {
ações: andar, buzinar e frear
*/
