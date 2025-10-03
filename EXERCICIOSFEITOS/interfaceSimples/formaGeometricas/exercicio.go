package main

import "fmt"

type Forma interface {
	Area() float64
}

type Quadrado struct {
	lado float64
}
type Retangulo struct {
	largura float64
	altura  float64
}

func (q Quadrado) Area() float64 {
	return q.lado * q.lado
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

func generica(f Forma) {
	fmt.Println(f.Area())
}

func main() {
	calculoQuadrado := Quadrado{lado: 4}
	calculoRetangulo := Retangulo{largura: 3, altura: 5}

	generica(calculoQuadrado)
	generica(calculoRetangulo)
}

/*Descrição:
Crie uma interface chamada Forma com um método:

Area() float64 → retorna a área da forma

Crie duas structs:

Quadrado → com atributo lado (float64)

Retangulo → com atributos largura e altura (float64)

Implemente Area() para cada struct.

No main, crie um quadrado e um retângulo e calcule suas áreas usando a interface.*/
