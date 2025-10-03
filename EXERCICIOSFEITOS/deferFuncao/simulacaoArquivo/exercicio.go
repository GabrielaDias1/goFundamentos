package main

import "fmt"

func processarArquivo() {

	fmt.Println("Abrindo arquivo")
	fmt.Println("Escrevendo arquivo")
	defer fmt.Println("Fechando arquivo")
	// Garante que o arquivo será fechado no final, independentemente do que acontecer
}

func main() {
	processarArquivo()
}

/*Crie uma função que abre um "arquivo" (simulado apenas com fmt.Println("Abrindo arquivo...")), depois faz alguma operação (fmt.Println("Escrevendo no arquivo...")), e use defer para garantir que o "arquivo seja fechado" (fmt.Println("Fechando arquivo...")) ao final da função.*/