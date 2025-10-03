package main

import "fmt"

type Comparador interface {
	Maior(a, b int) int
}

// criando interface atraves dos metodos para todos que implementarem o metodo ter o mesmo tipo

//criando struct
type Numero struct{}

// criando func implementando interface atraves do metodo
func (n Numero) Maior(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// func generica para lidar com todas as respostas juntas
func genericaInterface(c Comparador, a, b int) {
	resultado := c.Maior(a, b)

	fmt.Printf("O maior número entre %d e %d é: %d\n", a, b, resultado)

}
//sem necessidade de muitas linhas de codigo
//implementou o método Maior para a struct Numero
//A função genericaInterface não se importa qual tipo específico você passa, desde que esse tipo implemente a interface Comparador.

func main() {
	//criando objeto
	numeroAlea := Numero{}

	genericaInterface(numeroAlea, 10, 20)

}

/*Descrição:
Crie uma interface Comparador com o método:

Maior(a, b int) int → retorna o maior número entre dois.

Implemente em uma struct Numero.
Teste no main com dois valores diferentes.*/
