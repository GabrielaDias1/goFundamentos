package main

import "fmt"

type Operacao interface {
	Calcular(a, b int) int
}

type Soma struct{}
type Subtracao struct{}

func (s Soma) Calcular(a, b int) int {
	return a + b
}

func (s Subtracao) Calcular(a, b int) int {
	return a - b
}

// é Calcular(a, b int porque struct nao tem acesso a esses valores, a struct so sabe valores que esta nela

func executarOperacao(op Operacao, a, b int) {
	fmt.Println(op.Calcular(a, b))
}

func main() {

	// cria instancia

	executarOperacao(Soma{}, 2, 10)

	// nome para guardar algo que segue a regra da interface
	// (qualquer tipo que implemente o método Calcular pode ser usado aqui)
	// tipo genérico

	executarOperacao(Subtracao{}, 10, 5)

	// para reutlizar
	soma := Soma{} // cria e dá nome "soma"
	executarOperacao(soma, 10, 5)

}

/*Crie uma interface chamada Operacao com um método:

Calcular(a, b int) int → recebe dois números e retorna o resultado da operação
Crie duas structs:
Soma → implementa Calcular somando a + b
Subtracao → implementa Calcular subtraindo a - b
No main, use a interface para chamar Calcular de cada struct.*/
