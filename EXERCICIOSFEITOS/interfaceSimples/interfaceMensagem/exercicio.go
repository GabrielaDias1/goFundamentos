package main

import "fmt"

type Mensagem interface {
	Texto() string
}

type Email struct{}
type SMS struct{}

func (e Email) Texto() string {
	return "Mensagem enviada por e-mail"
}

func (s SMS) Texto() string {
	return "Mensagem enviada por SMS"
}

func generica(m Mensagem) {
	fmt.Println(m.Texto())
}

func main() {
	msgSms := SMS{}

	generica(msgSms)

	msgEmail := Email{}

	generica(msgEmail)
}

/*Descrição:
Crie uma interface chamada Mensagem com um método:

Texto() string → retorna o texto da mensagem

Crie duas structs:

Email → retorna "Mensagem enviada por e-mail"

SMS → retorna "Mensagem enviada por SMS"

Crie uma função que receba qualquer valor que implemente Mensagem e imprima o texto da mensagem.*/
