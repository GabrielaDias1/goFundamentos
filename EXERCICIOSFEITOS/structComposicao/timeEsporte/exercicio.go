package main

import "fmt"

type Jogador struct {
	Nome    string
	Posicao int
}

type Time struct {
	NomeTime  string
	Jogadores []Jogador
}

type Liga struct {
	NomeLiga string
	Times    []Time
}

func main() {
	ligaTitan := Liga{
		NomeLiga: "Liga Titan",
		Times: []Time{
			{
				NomeTime: "Time Liga Titan Volei",
				Jogadores: []Jogador{
					{Nome: "Julio", Posicao: 10},
					{Nome: "Julia", Posicao: 4},
					{Nome: "Amanda", Posicao: 2},
				},
			},
			{
				NomeTime: "Time Liga Titan Futebol",
				Jogadores: []Jogador{
					{Nome: "Gabriel", Posicao: 8},
					{Nome: "Gustavo", Posicao: 10},
					{Nome: "Jorge", Posicao: 2},
				},
			},
		},
	}

	fmt.Printf("Liga: %s\n", ligaTitan.NomeLiga)

	for _, time := range ligaTitan.Times {
		fmt.Printf("\nTime: %s\n", time.NomeTime)
		fmt.Println("Jogadores:")
		for _, jogador := range time.Jogadores {
			fmt.Printf(" - Nome: %s, Posição: %d\n", jogador.Nome, jogador.Posicao)
		}
	}
}

/*Time de esportes com jogadores

Crie uma struct Jogador com Nome e Posicao.

Crie uma struct Time com NomeTime e slice de Jogador.

Crie uma struct Liga com NomeLiga e slice de Time.

Crie uma liga com 2 times, cada um com 3 jogadores, e imprima os times e seus jogadores.*/
