package main

import "fmt"

type Carro struct {
	Modelo string
	Ano    int
}

func (c Carro) descricao() string {
	return fmt.Sprintf("Carro modelo %s do ano %d", c.Modelo, c.Ano)
}

func main() {
	carroCorolla := Carro{
		Modelo: "Corolla XEi 2.0 Flex",
		Ano:    2020,
	}

	carroHonda := Carro{
		Modelo: "Civic Touring 1.5 Turbo",
		Ano:    2019,
	}

	fmt.Println(carroCorolla.descricao())
	fmt.Println(carroHonda.descricao())

}

/*Método em struct

Crie uma struct Carro com Modelo e Ano.
Adicione um método chamado Descricao() que retorne uma string no formato:
"Carro modelo X do ano Y".
No main, crie dois carros e chame o método para exibir a descrição.*/
