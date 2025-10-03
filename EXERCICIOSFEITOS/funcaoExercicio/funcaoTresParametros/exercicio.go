package main

import "fmt"

func apresentacao(nome string, idade int, cidade string) string {

	msg := fmt.Sprintf("Olá, meu nome é %s, tenho %d anos e moro em %s.", nome, idade, cidade)
	return msg
}

func main() {
	fmt.Println(apresentacao("Gabriela", 22, "São Paulo"))
}

/*Função com três parâmetros

Crie uma função que recebe nome, idade e cidade e retorna uma frase formatada de apresentação.*/
