package main

import "fmt"

type Conversor interface {
	Converter(celsius float64) float64
}

type ParaFahrenheit struct{}

type ParaKelvin struct{}

func (f ParaFahrenheit) Converter(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func (k ParaKelvin) Converter(celsius float64) float64 {
	return celsius + 273.15
}

func genericaInterface(c Conversor, celsius float64) {
	resultado := c.Converter(celsius)
	fmt.Println(resultado)
}

func main() {
	numsCelsius := ParaFahrenheit{}

	numsCel := ParaKelvin{}

	genericaInterface(numsCelsius, 10.00)

	genericaInterface(numsCel, 10.00)
}

/*Descrição:
Crie uma interface Conversor com o método:

Converter(celsius float64) float64

Implemente em structs diferentes:

ParaFahrenheit

ParaKelvin

Teste no main com uma temperatura em Celsius.*/
