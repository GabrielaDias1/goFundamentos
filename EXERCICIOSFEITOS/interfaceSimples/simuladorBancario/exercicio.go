package main

import "fmt"

type Conta interface {
	Depositar(valor float64)
	Saldo() float64
}

type ContaCorrente struct {
	saldo float64
}

func (c ContaCorrente) Saldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Depositar(valor float64) {
	c.saldo += valor
}

func genericaInterface(c Conta, valor float64) {
	c.Depositar(valor)

	fmt.Println(c.Saldo())
}

func main() {
	contaNums := ContaCorrente{}

	genericaInterface(&contaNums, 400)
	genericaInterface(&contaNums, 100)
}

/*Crie uma interface Conta com os métodos:

Depositar(valor float64)

Saldo() float64

Implemente em uma struct ContaCorrente.
No main, deposite um valor e mostre o saldo.*/
