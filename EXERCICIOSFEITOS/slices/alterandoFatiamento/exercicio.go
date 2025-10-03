package main

import "fmt"

func main() {
	//Criação slice
	nums := []int{10, 20, 30, 40, 50}
	// criação segundo slice com fatia do primeiro que é nums
	parte := nums[1:4]
	// parte é [20, 30, 40, 50] fazendo com que indice 0 seja 20
	// Alterando o primeiro elemento da fatia
	parte[0] = 999

	fmt.Println("Slice original após alteração:", nums)
	fmt.Println("Fatia (parte) após alteração:", parte)

	/*saida:
	Slice original após alteração: [10 999 30 40 50]
	Fatia (parte) após alteração: [999 30 40]
	*/
	/*Observação:
	Slices compartilham o mesmo array interno.
	Alterar a fatia também altera o slice original, se a mudança afetar elementos compartilhados.
	*/
}

/*Alterando pelo fatiamento
Crie um slice nums := []int{10, 20, 30, 40, 50}.
Depois crie parte := nums[1:4] e mude o primeiro elemento de parte para 999.
imprima nums e parte.*/
