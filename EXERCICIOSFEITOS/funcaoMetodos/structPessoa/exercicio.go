package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}// metodo

func main() {
	gabriela := Pessoa{
		Nome:  "Gabriela",
		Idade: 22,
	}// instancia/objeto

	gabriela.apresentar()
}
