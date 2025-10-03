package main

import "fmt"

func exibirElementos(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
	//o Go entende que você pode passar quantos números quiser, e todos eles vão ser agrupados dentro de um slice de inteiros chamado x.

}

func main() {
	numero := []int{1, 2, 3, 4, 5}

	exibirElementos(numero...)
}

/*Crie uma função que receba um slice de números inteiros e mostre todos os elementos.*/
