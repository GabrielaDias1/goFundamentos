package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("todos os numeros: ", numeros)
	//pegar todos os elementos do início do slice até, mas não incluindo, o índice 2
	//pegar todos os elementos a partir do índice 3 até o fim do slice
	numeros = append(numeros[:2], numeros[3:]...) //... é para seguir apos o 3, seguindo para os proximos numeros, desenrola a partir do indice 3, quer dizer que venha os numeros a partir do indice 3 
	fmt.Println("numeros apos append, ", numeros)// essa fatia exclui o tres
}

/*Removendo elemento do meio
Crie numeros := []int{1, 2, 3, 4, 5} e remova o número 3 usando:
numeros = append(numeros[:2], numeros[3:]...)
Imprima antes e depois.*/
