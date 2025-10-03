package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	//cria dois slice

	c := append(a, b...)
	//pegue todos os elementos de b e passe um por um para o append

	fmt.Println(c)
}

/*Crie dois slices: a := []int{1, 2} e b := []int{3, 4}.
Junte os dois em um novo slice c usando append.
Imprima o resultado.*/
