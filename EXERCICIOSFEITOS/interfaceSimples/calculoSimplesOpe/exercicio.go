package main

import "fmt"

//cria a interface com o metodo calcular
type Operacao interface {
	Calcular(a, b int) int
}

type Soma struct{}
type Multiplicacao struct{}

func (s Soma) Calcular(a, b int) int {
	return a + b
}

func (m Multiplicacao) Calcular(a, b int) int {
	return a * b
}

func genericaInterface(o Operacao, a, b int) {
	resultado := o.Calcular(a, b)
	fmt.Println("Resultado:", resultado)
}

func main() {
	numsAleatorio := Soma{}
	genericaInterface(numsAleatorio, 10, 10)

	numeroAleatorio := Multiplicacao{}
	genericaInterface(numeroAleatorio, 20, 20)
}

/*Descrição:
Crie uma interface Operacao com um método:

Calcular(a, b int) int

Implemente em structs diferentes:

Soma

Multiplicacao

No main, use a interface para executar uma operação escolhida.*/
