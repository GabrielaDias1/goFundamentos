package main

import "fmt"

func main() {
	// Slice com números de 1 a 10
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Fatiando os 3 primeiros números (índices 0, 1, 2) exclui 3
	primeirosNumeros := numeros[0:3]
	// Fatiando os 3 últimos números (índices 7, 8, 9) exclui 10
	ultimosNumeros := numeros[7:10]

	fmt.Println("Todos os numeros: ", numeros)
	fmt.Println("Fatia dos primeiros numeros: ", primeirosNumeros)
	fmt.Println("Fatia dos ultimos numeros: ", ultimosNumeros)
}

/*Fatiamento básico
Crie um slice de números de 1 a 10 e crie:

Um slice com os 3 primeiros números

Um slice com os 3 últimos números
Imprima cada um.

*/
