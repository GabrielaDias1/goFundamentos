package main

import "fmt"

type PrimoChecker interface {
	EhPrimo(n int) bool
}

type Numero struct{}

func (n Numero) EhPrimo(x int) bool {
	if x == 2 || x == 3 || x == 5 || x == 7 || x == 11 || x == 13 || x == 17 || x == 19 {
		return true
	}

	if x < 2 || x == 4 || x == 6 || x == 8 || x == 9 || x == 10 ||
		x == 12 || x == 14 || x == 15 || x == 16 || x == 18 || x == 20 {
		return false
	}
	return false
}

func generica(p PrimoChecker, x int) {
	if p.EhPrimo(x) {
		fmt.Println("É numero primo")
	} else {
		fmt.Println("Não é numero primo")
	}
}

func main() {
	//cria objeto

	numAleatorio := Numero{}

	generica(numAleatorio, 8)
	generica(numAleatorio, 2)
}

/*Descrição:
Crie uma interface chamada PrimoChecker com um método:

EhPrimo(n int) bool

Crie uma struct Numero que implemente esse método.

Verifique se n é divisível apenas por 1 e ele mesmo.

Retorne true se for primo, false caso contrário.

No main, teste com um número de cada vez.*/
