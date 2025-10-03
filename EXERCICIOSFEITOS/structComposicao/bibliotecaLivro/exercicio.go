package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Ano    int
}

type Biblioteca struct {
	Nome  string
	Livro []Livro
}

func main() {
	bibliotecaPublica := Biblioteca{
		Nome: "Biblioteca publica",
		Livro: []Livro{
			{Titulo: "harry potter e a pedra filosofal", Autor: "J. K. Rowling", Ano: 2001},
			{Titulo: "percy jackson e o ladrão de raios", Autor: "Rick Riordan", Ano: 2005},
			{Titulo: "senhor dos aneis a Sociedade do Anel", Autor: "J. R. R. Tolkien", Ano: 1954},
		},
	}
	fmt.Printf("A %s\n", bibliotecaPublica.Nome)
	fmt.Println("Livros que essa biblioteca possui: ")

	for _, l := range bibliotecaPublica.Livro {
		fmt.Printf("Titulo: %s | Autor: %s | Ano lançamento: %d\n", l.Titulo, l.Autor, l.Ano)
	}

}

/*Biblioteca com livros

Crie uma struct Livro com Titulo, Autor e Ano.

Crie uma struct Biblioteca com Nome e uma slice de Livro.

Adicione 3 livros à biblioteca e imprima:
"Biblioteca X tem os livros: L1, L2, L3".*/
