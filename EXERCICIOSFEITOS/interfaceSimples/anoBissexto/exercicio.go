package main

import "fmt"

type Calendario interface {
	EhBissexto(ano int) bool
}

type Ano struct{}

func (n Ano) EhBissexto(ano int) bool {
	if (ano%4 == 0) && ((ano%100 != 0) || (ano%400 == 0)) {
		return true
	} else {
		return false
	}
}

func genericoInterface(c Calendario, ano int) {
	resultado := c.EhBissexto(ano)

	fmt.Println("O ano é bissexto!", resultado)
}

func main() {
	numsAno := Ano{}

	genericoInterface(numsAno, 2025)
	genericoInterface(numsAno, 2024)
	genericoInterface(numsAno, 2020)
}

/*Descrição:
Crie uma interface Calendario com o método:

EhBissexto(ano int) bool

Implemente na struct Ano.
Regras:

Divisível por 4 → bissexto

Exceto se divisível por 100, a menos que seja divisível por 400

Teste no main.*/
