package main

import "fmt"

type Checagem interface {
	EhPar(n int) bool
}

type Numero struct{}

func (e Numero) EhPar(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func generica(c Checagem, n int) {
	fmt.Printf("%d é par? %v\n", n, c.EhPar(n))
}

func main() {
	var c Checagem = Numero{}

	numeros := []int{1, 2, 3, 10, 15, 20}

	for _, n := range numeros {
		generica(c, n)
	}
}

/*Descrição:
Crie uma interface chamada Checagem com um método:

EhPar(n int) bool → retorna true se o número for par e false caso contrário

Crie uma struct chamada Numero.

No main, use a interface para verificar se vários números são pares ou ímpares.*/
