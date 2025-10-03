package main

import "fmt"

func main() {
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i, linha := range matriz {
		for j, valor := range linha {
			fmt.Printf("matriz[%d][%d] = %d\n", i, j, valor)
		}
	}
	/*linha = a linha inteira (um slice).
	valor = o número específico dentro daquela linha.*/
}

/*Crie um slice bidimensional ([][]int) que represente uma matriz de 2 linhas e 3 colunas:*/
