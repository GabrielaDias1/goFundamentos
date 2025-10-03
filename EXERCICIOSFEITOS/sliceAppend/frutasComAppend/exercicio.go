package main

import "fmt"

func main() {
	frutas := []string{}
	var fruta string

	for {
		fmt.Print("Digite nome da fruta (ou 'fim' para parar): ")
		fmt.Scan(&fruta)

		if fruta == "fim" {
			break
		}

		frutas = append(frutas, fruta)
	}

	fmt.Println("Frutas digitadas:", frutas)
}

/*Lista de frutas

Crie um slice vazio de string e peça para o usuário digitar nomes de frutas.
Adicione cada fruta com append até que o usuário digite "fim".
Depois, imprima todas as frutas.*/
