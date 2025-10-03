package main

import "fmt"

func main() {
	cidades := make([]string, 2, 4)
	//crie o slice e inseri o make para ter tamanho e capacidade pre-definido
	// chamo o slice coloco seu tipo, tamanho e capacidade
	cidades[0], cidades[1] = "são paulo", "rio de janeiro"
	//inicializando ja com valores
	fmt.Println(cidades)
	//imprimir antes do apppend

	cidades = append(cidades, "salvador", "brasilia")
	fmt.Println(cidades)
	//[são paulo rio de janeiro] -> correto

}

/*Crie um slice de strings com make (tamanho 2, capacidade 4).
➝ Preencha com 2 nomes de cidades.
➝ Adicione mais 2 cidades com append.*/
