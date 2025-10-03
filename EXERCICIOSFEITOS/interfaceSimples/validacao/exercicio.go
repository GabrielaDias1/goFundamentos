package main

import "fmt"

type Validador interface {
	Valido() bool
}

type ContaBancaria struct {
	saldo float64
}

func (c ContaBancaria) Valido() bool {
	if c.saldo > 0 {
		return true
	} else {
		return false
	}
}

func generica(v Validador) {
	fmt.Println(v.Valido())
}

func main() {
	conta1 := ContaBancaria{saldo: 100.00}
	conta2 := ContaBancaria{saldo: -50.00}
	conta3 := ContaBancaria{saldo: 0.00}

	generica(conta1)
	generica(conta2)
	generica(conta3)

}

/*Descrição:
Crie uma interface chamada Validador com um método:

Valido() bool → retorna true ou false dependendo da validação

Crie uma struct ContaBancaria com um atributo:

saldo (float64)

Implemente o método Valido() para retornar true se o saldo for maior que zero e false caso contrário.

No main, crie algumas contas e verifique se são válidas usando a interface.*/
