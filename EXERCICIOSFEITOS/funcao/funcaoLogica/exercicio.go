package main

import "fmt"

func main() {

	fmt.Println(ehPar(5))
}

func ehPar(numero int) bool {
	if numero%2 == 0 {
		return true
	}
	return true
}

/*Função com lógica simples

Crie uma função chamada ehPar(n int) bool que retorna true se o número for par e false se for ímpar.*/
