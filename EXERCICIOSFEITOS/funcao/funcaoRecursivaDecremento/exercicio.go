package main

import "fmt"

func contagem(numero int) {

	if numero == 0 {
		return 
	}
	fmt.Println("Número:", numero)
	contagem(numero - 1) 
}

func main() {
	contagem(5)
}

/*Função recursiva simples

Crie uma função contagem(n int) que imprime os números de n até 1 usando recursão.*/
