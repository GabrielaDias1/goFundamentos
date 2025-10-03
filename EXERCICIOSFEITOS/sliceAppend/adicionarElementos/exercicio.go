package main

import "fmt"

func main() {
	frutas := []string{"maçã", "banana", "caju"}
	// Criado o slice: ele não precisa de tamanho fixo. Um slice é dinâmico — não é preciso declarar o tamanho inicial.

	frutas = append(frutas, "morango", "melancia")
	// append adiciona novos elementos ao final do slice

	fmt.Println(frutas)
}

//Crie um slice de strings com 3 frutas e adicione mais 2 usando append.
