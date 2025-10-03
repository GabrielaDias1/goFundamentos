// Crie uma variável dia com o nome de um dia da semana. Use switch para imprimir uma mensagem diferente para "segunda", "quarta" e "sexta". Caso não seja nenhum desses dias, imprima "Outro dia".
package main

import "fmt"

func main() {
	var dia = "sexta"

	switch dia {
	case "segunda":
		fmt.Println("Boa segunda-feira!")
	case "quarta":
		fmt.Println("Bem vindo!")
	case "sexta":
		fmt.Println("Bom final de semana")
	default:
		fmt.Println("Outro dia")
	}

}

/*Crie uma variável dia com o nome de um dia da semana. Use switch para imprimir uma mensagem diferente
para "segunda", "quarta" e "sexta". Caso não seja nenhum desses dias, imprima "Outro dia".*/


















