package main

import "fmt"

func main() {
	numeros := []int{2, 4, 6}
	// cria um slice, slice sempre tamanho vazio
	fmt.Println(numeros)
	// aqui teste antes de usar o append
	numeros = append(numeros, 8, 10)
	// usando append para adiciona numeros no final 
	fmt.Println(numeros)
	// para teste imprimir apos append [2 4 6 8 10] -> correto

}

/*Crie um slice de números [2, 4, 6].
➝ Adicione os números 8 e 10.*/
