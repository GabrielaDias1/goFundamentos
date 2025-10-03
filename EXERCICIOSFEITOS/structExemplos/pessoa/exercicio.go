package main

import "fmt"

type Pessoa struct {
	Nome   string
	altura string
	Etnia  string
	Idade  int
}

func (p Pessoa) trabalhar() {
	fmt.Printf("%s, com %s de altura, %d anos e etnia %s, está trabalhando\n", p.Nome, p.altura, p.Idade, p.Etnia)

}

func (p Pessoa) cafe() {
	fmt.Printf("A %s está tomando café.\n", p.Nome)
}

func (p Pessoa) assistir() {
	fmt.Printf("A %s esta assistindo Tv\n", p.Nome)
}

func (p Pessoa) fofoca() {
	fmt.Printf("Agora a %s esta fofocando\n", p.Nome)
}
func main() {
	anaBeatriz := Pessoa{
		Nome:   "Ana Beatriz",
		altura: "1,65m",
		Etnia:  "parda",
		Idade:  28,
	}
	anaBeatriz.trabalhar()
	anaBeatriz.cafe()
	anaBeatriz.assistir()
	anaBeatriz.fofoca()

}

/*struct pessoa

propriedades: etnia, idade, nome, altura

metodos: trabalhar, tomar cafe, assistir, fofoca,*/
