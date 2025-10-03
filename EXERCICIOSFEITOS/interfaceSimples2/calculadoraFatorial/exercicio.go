package main

import "fmt"

type Calculo interface {
	Fatorial(n int) int
}

type Numero struct{}

func (e Numero) Fatorial(n int) int {
	fatorial := 1
	for i := 1; i <= n; i++ {
		fatorial = fatorial * i
	}
	return fatorial
}

func genericaInterface(c Calculo, n int) {
	resultado := c.Fatorial(n)

	fmt.Printf("Fatorial do número %d é %d\n", n, resultado)
}
func main() {
	numsFatorial := Numero{}

	genericaInterface(numsFatorial,5)
}

/*Descrição:
Crie uma interface Calculo com o método:

Fatorial(n int) int

Implemente em uma struct Numero.
Teste no main para 5! = 120.*/
