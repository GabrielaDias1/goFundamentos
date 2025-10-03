package main

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func (r Retangulo) calculoArea() float64 {

	area := r.largura * r.altura
	return area
}

func main() {
	retangulo1 := Retangulo{
		largura: 4.0,
		altura:  5.0,
	}

	fmt.Println(retangulo1.calculoArea())// Saída: 20
}
