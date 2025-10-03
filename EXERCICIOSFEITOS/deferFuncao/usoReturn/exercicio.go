package main

import "fmt"

func soma(a int, b int) int {
	defer fmt.Println("Finalizando função")// sera executado antes da funcao acabar

	return a + b
	//defer apos return deu erro 

}

func main() {
	fmt.Println(soma(1, 1))
}
/*Escreva uma função que retorna um valor (int ou string) e dentro dela coloque um defer fmt.Println("Finalizando função").
Teste para ver se o defer é executado mesmo antes do retorno*/