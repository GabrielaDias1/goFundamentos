package main

import "fmt"

type Carro struct {
	Nome string
	Cor  string
	Ano  int
}

func (c Carro) andar() {
	fmt.Printf("O carro %s, da cor %s e do ano %d, está andando.\n", c.Nome, c.Cor, c.Ano)
}

func main() {
	gol := Carro{
		Nome: "gol",
		Cor:  "preto",
		Ano:  2000,
	}

	gol.andar()
}
