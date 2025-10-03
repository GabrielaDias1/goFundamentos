package main

import (
	"fmt"
	"strings"
)

type Filtro interface {
	ContemPalavra(s string) bool
}

type Moderador struct{}

func (m Moderador) ContemPalavra(s string) bool {
	s = strings.ToLower(s)
	return strings.Contains(s, "proibido")
	//serve para verificar se a string s contém a palavra "proibido".

}

func genericaInterface(f Filtro, s string) {
	resultado := f.ContemPalavra(s)
	fmt.Println(resultado)
}

func main() {
	moderador := Moderador{}

	frase1 := "Este conteúdo é proibido para menores."
	frase2 := "Este conteúdo está liberado."

	genericaInterface(moderador, frase1)
	genericaInterface(moderador, frase2)
}

/*Descrição:
Crie uma interface Filtro com o método:

ContemPalavra(s string) bool

Implemente em uma struct Moderador.

Retorne true se a string contiver "proibido"

Caso contrário, false.

Teste no main com duas frases.*/
