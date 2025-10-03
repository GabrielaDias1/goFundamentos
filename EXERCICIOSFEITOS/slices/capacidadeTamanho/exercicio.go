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

	fmt.Println("Slice original:", nums)
	fmt.Println("Slice original tamanho:", len(nums))    // quantidade de elementos que o slice possui
	fmt.Println("Slice original capacidade:", cap(nums)) ////capacidade máxima antes de precisar criar um novo array interno
	fmt.Println("Fatia (parte) após alteração:", parte)
	fmt.Println("Fatia (parte) após alteração tamanho:", len(parte))//len(parte) → quantos elementos você pegou → 3 ([999, 30, 40])
	fmt.Println("Fatia (parte) após alteração capacidade:", cap(parte))//quantos elementos ainda cabem no array interno a partir do índice inicial da fatia → 4 ([999, 30, 40, 50])

	/*saida:
	Slice original após alteração: [10 999 30 40 50]
	Fatia (parte) após alteração: [999 30 40]
	*/
	/*Observação:
	Slices compartilham o mesmo array interno.
	Alterar a fatia também altera o slice original, se a mudança afetar elementos compartilhados.
	*/
	/*cap(parte) = total de elementos do array interno
        a partir do índice inicial da fatia
		
*/
}

/*Capacidade e tamanho
Usando o slice nums do exercício anterior, imprima len(nums), cap(nums), len(parte) e cap(parte).*/
