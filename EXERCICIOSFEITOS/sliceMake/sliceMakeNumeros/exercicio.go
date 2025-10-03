package main

import "fmt"

func main() {
	numeros := make([]int, 0, 5)
	//len -> quantos elementos tem
	//cap -> sua capacidade
	// Criando um slice de int com make: tamanho 0, capacidade 5

	numeros = append(numeros, 1, 2, 3, 4, 5)
	fmt.Println("Slice:", numeros)

	// Mostrando tamanho (len) e capacidade (cap)
	fmt.Println("Tamanho (len):", len(numeros))
	//retorna quantos elementos tem.
	fmt.Println("Capacidade (cap):", cap(numeros))
	//retorna a capacidade máxima antes de expandir
}

/*Crie um slice de int usando make([]int, 0, 5).
Adicione os números 1, 2, 3, 4, 5 com append.
Mostre o tamanho e a capacidade do slice no final.*/
