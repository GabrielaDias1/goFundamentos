package main

import "fmt"

type Pessoa struct {
	// cria a struct pessoa
	Nome  string
	Idade int
	//A struct agrupa os campos Nome e Idade, permitindo criar variáveis do tipo Pessoa
}

func main() {

	// aqui é possivel manipular a struct pessoa com informações, é uma variavel que ja é do tipo pessoa
	pessoa1 := Pessoa{
		Nome:  "gaby",
		Idade: 22,
	}
	pessoa2 := Pessoa{
		Nome:  "elias",
		Idade: 70,
	}

	fmt.Println("Primeira pessoa: ", pessoa1)
	fmt.Println("Segunda pessoa: ", pessoa2)

}

/*Cadastro simples
Crie uma struct Pessoa com os campos Nome (string) e Idade (int).
No main, crie duas pessoas e imprima no console.*/
