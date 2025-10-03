package main

import "fmt"

func main() {
	numeros := make([]int, 3, 5)
	// crie um slice numeros e inseri o make para ter cap e tamanho pre-definido
	//len(numeros) (tamanho) = 3 → já tem 3 posições reservadas.
	//cap(numeros) (capacidade) = 5 → pode crescer até 5 sem precisar realocar memória.
	numeros[0], numeros[1], numeros[2] = 1, 2, 3
	fmt.Println(numeros)
	//imprimir para ver antes do append
	numeros = append(numeros, 4, 5)
	fmt.Println(numeros)

}

/*Use make para criar um slice de int com tamanho 3 e capacidade 5.
➝ Preencha os 3 primeiros elementos com números à sua escolha.
➝ Depois, use append para adicionar mais 2 números e verifique o slice.*/
