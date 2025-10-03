package main

import "fmt"

func main() {
	grupos := [][]string{
		{"ana", "bia"},
		{"carol", "duda"},
	}
	// Adicionando um terceiro grupo
	grupos = append(grupos, []string{"fernanda", "gabi"})
	fmt.Println(grupos)
}

/*Grupos de estudo

Crie um slice [][]string onde cada grupo tem 2 alunos.
Adicione 3 grupos com append.
Depois, imprima os grupos um por um.*/
