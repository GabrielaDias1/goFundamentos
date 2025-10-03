package main

import "fmt"

func main() {
	frutas := []string{"ana", "bia", "carol", "duda", "eli"}
	fmt.Println("seu tamanho", len(frutas))    // quantidade de elementos que o slice possui
	fmt.Println("sua capacidade", cap(frutas)) //capacidade máxima antes de precisar criar um novo array interno
}

/*Crie um slice com 5 nomes de frutas e imprima o slice, seu tamanho (len) e capacidade (cap).*/
