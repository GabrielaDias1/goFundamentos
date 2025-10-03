package main

import "fmt"

type Carro struct {
	Modelo string
	Ano    int
	Marca  string
}

func main() {
	ListaCarro := []Carro{
		{Modelo: "Camaro", Ano: 2019, Marca: "Chevrolet"},
		{Modelo: "Corolla", Ano: 2020, Marca: "Toyota"},
		{Modelo: "Wrangler", Ano: 2022, Marca: "Jeep"},
	}

	for _, m := range ListaCarro {
		fmt.Printf("O carro %s do ano %d é da marca %s\n", m.Modelo, m.Ano, m.Marca)
	}
}

/*Lista de carros

Crie uma slice de Carro e adicione 3 carros.
Use um for para imprimir cada carro com uma frase:
"O carro X do ano Y é da marca Z".*/
