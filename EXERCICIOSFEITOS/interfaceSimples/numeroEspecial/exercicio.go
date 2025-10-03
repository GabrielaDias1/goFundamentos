package main

import "fmt"

type NumeroEspecial interface {
	EhPerfeito(n int) bool
}

type numero struct{}

func (x numero) EhPerfeito(n int) bool {
	soma := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			soma = soma + i

		}

	}
	if soma == n {
		return true
	} else {
		return false
	}
}

func genericaInterface(y NumeroEspecial, n int) {
	resultado := y.EhPerfeito(n)

	fmt.Println(resultado)
}

func main() {
	numsAleatorio := numero{}

	genericaInterface(numsAleatorio, 496)
}

/*Descrição:
Crie uma interface NumeroEspecial com o método:

EhPerfeito(n int) bool

Um número é perfeito se a soma de seus divisores (exceto ele mesmo) é igual ao número.
Ex: 6 → divisores (1, 2, 3), soma = 6.

Implemente na struct Numero e teste no main.*/
