package main

import "fmt"

type Gato struct {
	Nome string
	Cor  string
	Raca string
}

func (g Gato) andar() {
	fmt.Printf("O gato %s, da cor %s e da raça %s, está andando.\n", g.Nome, g.Cor, g.Raca)
}

func (g Gato) miar() {
	fmt.Printf("Miau miau!\n")
}

func (g Gato) bebendoAgua() {
	fmt.Printf("O gato %s está bebendo agua\n", g.Nome)
}

func main() {
	luna := Gato{
		Nome: "Luna",
		Cor:  "Cinza",
		Raca: "Gato Siberiano",
	}

	luna.andar()
	luna.miar()
	luna.bebendoAgua()

}

/*struct gato:

propriedade: nome, cor, raça

metodos : andar, miar, beberAgua*/
