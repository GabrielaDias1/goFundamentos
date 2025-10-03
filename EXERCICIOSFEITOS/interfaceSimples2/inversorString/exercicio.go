package main

import "fmt"

type Texto interface {
	Inverter(s string) string
}

type Mensagem struct{}

func (m Mensagem) Inverter(s string) string {

	resultado := ""

	for i := len(s) - 1; i >= 0; i-- {
		resultado = resultado + string(s[i])
	}

	return resultado

}

func genericaInterface(t Texto, s string) {
	invertido := t.Inverter(s)

	fmt.Printf("Resultado palavra %s invertida %s ", s, invertido)
}

func main() {
	palavraAleatoria:=Mensagem{}

	genericaInterface(palavraAleatoria,"golang")
}

/*Descrição:
Crie uma interface Texto com o método:

Inverter(s string) string

Implemente em uma struct Mensagem.
Exemplo: "golang" → "gnalog".*/