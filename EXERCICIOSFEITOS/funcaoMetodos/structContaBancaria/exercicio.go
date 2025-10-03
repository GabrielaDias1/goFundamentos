package main

import "fmt"

type contaBancaria struct {
	titular string
	saldo   float64
}

func (c contaBancaria) sacar(valorSaque float64) (float64, string) {
	if c.saldo <= 0 {
		return c.saldo, "Saque Insuficiente"
	} else if c.saldo >= valorSaque {
		c.saldo -= valorSaque
		return c.saldo, "Saque realizado com sucesso"
	} else {
		return c.saldo, "Saque indisponivel"
	}

}

func main() {

	contaGabriela := contaBancaria{
		titular: "Gabriela",
		saldo:   500.00,
	}
	fmt.Println(contaGabriela.sacar(100.00))
}

/*Crie uma struct ContaBancaria com titular e saldo.
Faça um método Sacar(valor float64) que diminui o saldo, mas só se houver dinheiro suficiente.*/
