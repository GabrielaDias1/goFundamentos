package main

import (
	"fmt"
	"strings"
)

type AnaliseTexto interface {
	TemVogal(s string) bool
}

type Frase struct{}

func (f Frase) TemVogal(s string) bool {
	letra := strings.ToLower(s)

	for i := len(letra) - 1; i >= 0; i-- {
		if letra[i] == 'a' || letra[i] == 'e' || letra[i] == 'i' || letra[i] == 'o' || letra[i] == 'u' {
			return true
		}
	}
	return false
}
func genericaInterface(a AnaliseTexto, s string) {
	resultado := a.TemVogal(s)

	fmt.Printf("Tem vogal? %t\n", resultado)

}

func main() {
	palavraAleatoria := Frase{}

	genericaInterface(palavraAleatoria, "casa")
	genericaInterface(palavraAleatoria, "rhythm")
}

/*Descrição:
Crie uma interface AnaliseTexto com o método:

TemVogal(s string) bool

Implemente em uma struct Frase.
No main, teste com "casa" (true) e "rhythm" (false).*/
