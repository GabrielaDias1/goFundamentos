package main

import "fmt"

func verificaPar(numero int) string {
	if numero%2 == 0 {
		return "par"
	}
	return "impar"
}

func main() {
	fmt.Println(verificaPar(4))
	fmt.Println(verificaPar(5))

}

/*Função com um parâmetro

Crie uma função que recebe um número inteiro e retorna se ele é par ou ímpar.*/
