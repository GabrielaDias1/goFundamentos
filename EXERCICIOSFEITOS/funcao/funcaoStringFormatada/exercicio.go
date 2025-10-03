package main

import "fmt"

func main() {
	apresentar("Gabriela", 22)
}

func apresentar(nome string, idade int) string {
	msg := fmt.Sprintf("Meu nome é %s e tenho %d anos", nome, idade)
	return msg
}

/*Crie uma função apresentar(nome string, idade int) string que retorna uma frase no formato:
"Meu nome é [nome] e tenho [idade] anos."*/
