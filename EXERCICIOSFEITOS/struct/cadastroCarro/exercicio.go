package main

import "fmt"

type Carro struct {
	Marca  string
	Modelo string
	Ano    int64
}

func main() {
	//// Criação de dois carros
	carro1 := Carro{
		Marca:  "Ford",
		Modelo: "Mustang",
		Ano:    2021,
	}
	carro2 := Carro{
		Marca:  "Volkswagen",
		Modelo: "Golf",
		Ano:    2018,
	}
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", carro1.Marca, carro1.Modelo, carro1.Ano)
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", carro2.Marca, carro2.Modelo, carro2.Ano)

}

/*Cadastro de carro

Crie uma struct Carro com os campos Marca (string), Modelo (string) e Ano (int).
Crie 2 carros diferentes e imprima no console no formato:
"Marca: X, Modelo: Y, Ano: Z".*/
