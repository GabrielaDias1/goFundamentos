package main

import "fmt"

type VerificadorIdade interface {
	MaiorDeIdade(idade int) bool
}

type Pessoa struct{}

func (p Pessoa) MaiorDeIdade(idade int) bool {
	if idade >= 18 {
		return true
	} else {
		return false
	}
}

func generica(v VerificadorIdade, idade int) {
	if v.MaiorDeIdade(idade) {
		fmt.Printf("é maior de idade!")
	} else {
		fmt.Println("não é maior de idade")
	}
}

func main() {
	gabriela := Pessoa{}

	generica(gabriela, 22)


}

/*Descrição:
Crie uma interface chamada VerificadorIdade com um método:

MaiorDeIdade(idade int) bool

Crie uma struct Pessoa que implemente esse método.

Retorne true se a idade for maior ou igual a 18.

Retorne false caso contrário.

No main, teste com apenas uma idade por vez.*/
