package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Ano    int
}

func main() {
	livro1 := Livro{
		Titulo: "Percy Jackson e o Ladrão de Raios",
		Autor:  "Rick Riordan",
		Ano:    2005,
	}
	livro2 := Livro{
		Titulo: "Harry Potter e a Pedra Filosofal",
		Autor:  "J.K. Rowling",
		Ano:    1997,
	}
	fmt.Println("Livros", livro1, livro2)

}

/*Cadastro de livro

Crie uma struct Livro com os campos Titulo (string), Autor (string) e Ano (int).
No main, crie dois livros e imprima cada um no console.*/
