package main

import "fmt"

func main() {
	numeros := [3]int{1, 2, 3}
	fmt.Println("antes da alteração:", numeros)
	alterar(numeros)//Aqui o array é copiado e a função recebe uma cópia.
	fmt.Println("depois da alteração:", numeros)
}

func alterar(nums [3]int)  {
	nums[0] = 999
	fmt.Println("dentro da função alterar:", nums)
}

/*Passagem por valor
Faça uma função que receba um array de 3 inteiros e modifique o primeiro elemento para 999 dentro da função. No main, imprima o array antes e depois da chamada da função.*/
