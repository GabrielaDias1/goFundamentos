package main

import "fmt"

type Analisador interface {
	EhPositivo(n int) bool
}

type Numero struct{}

func (e Numero) EhPositivo(n int) bool {
	if n >= 0 {
		return true
	} else {
		return false
	}
}

func generica(a Analisador, n int) {
	if a.EhPositivo(n) {
		fmt.Printf("O número %d é positivo.\n", n)
	} else {
		fmt.Printf("O número %d é negativo.\n", n)
	}
}

func main() {

	teste1 := Numero{}
	generica(teste1, 2)

	teste2 := Numero{}

	generica(teste2, -5)
}

/*Positivo ou Negativo

Descrição:
Crie uma interface chamada Analisador com um método:

EhPositivo(n int) bool

Crie uma struct Numero que implemente esse método.

Retorne true se o número for maior ou igual a 0.

Retorne false caso contrário.

No main, teste com apenas um número de cada vez.*/
