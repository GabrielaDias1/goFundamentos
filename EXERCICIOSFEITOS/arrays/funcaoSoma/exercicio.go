package main

import "fmt"

// Função principal, ponto de entrada do programa
func main() {
	// Cria um array chamado 'numeros' com 4 elementos inteiros
	numeros := [4]int{1, 2, 3, 4}

	// Chama a função Soma passando o array 'numeros'
	// e imprime o resultado junto com o texto "soma: "
	fmt.Println("soma: ", Soma(numeros))
}

// Função que recebe um array de 4 inteiros e retorna a soma dos elementos
func Soma(nums [4]int) int {
	// Inicializa uma variável 'soma' com zero para acumular os valores
	soma := 0

	// Percorre cada elemento do array 'nums'
	// '_' ignora o índice e 'valor' recebe o valor atual
	for _, valor := range nums {
		// Adiciona o valor atual à variável 'soma'
		soma += valor
	}

	// Retorna o total acumulado na variável 'soma'
	return soma
}

/*Função para soma
Escreva uma função que receba um array de 4 inteiros e retorne a soma de seus elementos. Chame essa função no main e imprima o resultado.*/
