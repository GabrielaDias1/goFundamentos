package main

import "fmt"

type Saudacao interface {
	Mensagem(nome string) string
}

type SaudacaoFormal struct{}
type SaudacaoInformal struct{}

func (f SaudacaoFormal) Mensagem(nome string) string {
	msg := fmt.Sprintf("Boa noite, Sr(a) %s", nome)
	return msg
}

func (i SaudacaoInformal) Mensagem(nome string) string {
	mens := fmt.Sprintf("Oi, %s!", nome)
	return mens
}

func genericaInterface(s Saudacao, nome string) {
	resultadoMsg := s.Mensagem(nome)

	fmt.Println(resultadoMsg)
}

func main() {
	msgFormal := SaudacaoFormal{}
	msgInformal := SaudacaoInformal{}

	genericaInterface(msgFormal, "Gabriela")

	genericaInterface(msgInformal, "Eliana")
}

/*Gerador de Saudações

Descrição:
Crie uma interface Saudacao com o método:

Mensagem(nome string) string

Implemente em duas structs:

SaudacaoFormal → "Boa noite, Sr(a). X"

SaudacaoInformal → "Oi, X!"

No main, mostre as duas saudações.*/
