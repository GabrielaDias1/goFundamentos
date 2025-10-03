package main

import "fmt"

func media(a float64, b float64) float64 {
	media := (a + b) / 2
	return media
}

func main() {
	fmt.Println(media(8.5, 7.0))
}

/*Função de cálculo de média

Crie uma função media(a float64, b float64) float64 que recebe duas notas e retorna a média delas.*/
