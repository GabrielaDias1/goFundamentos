package main

import "fmt"

func main() {
	defer fmt.Println("1") //ultimo 
	defer fmt.Println("2") // segundo 
	defer fmt.Println("3") // primeiro 
}
//O último que entra é o primeiro que sai

//Implemente 3 defer seguidos que imprimam mensagens como "Primeiro", "Segundo", "Terceiro".//