package main

import "fmt"

type Exibivel interface {
	Exibir() string
}

// criando interface e metodo

type Aluno struct {
	nome string
	nota int
}

// criando struct e seus atributos

type Professor struct {
	nome       string
	disciplina string
}

// implementando metodo na struct
func (a Aluno) Exibir() string {
	return fmt.Sprintf("%s - Nota: %d", a.nome, a.nota)
	// pensando em coloca sprint por ter apenas um return string
}

func (p Professor) Exibir() string {
	return fmt.Sprintf("%s - Disciplina: %s", p.nome, p.disciplina)
}

//quero criar a funcao generica para ja mostrar as implementacoes da interfaces nos metodos existentes
// pensando que na gererica devo pensa em como vou os objetos criados a partir então quando chama isso vai ter que refletir
func generica(e []Exibivel) {

	for _, v := range e {
		fmt.Println(v.Exibir())
	}
	//fmt.Println(e.Exibir())
}

func main() {
	// Criando slice de Exibivel direto, com instâncias anônimas
	exibiveis := []Exibivel{
		Aluno{nome: "Gabriela", nota: 10},
		Aluno{nome: "Lucas", nota: 8},
		Professor{nome: "Amanda", disciplina: "Matemática"},
		Professor{nome: "Carlos", disciplina: "História"},
	}
	generica(exibiveis)
}

/*Descrição:
Crie uma interface chamada Exibivel com um método:

Exibir() string → retorna uma string que descreve o objeto

Crie duas structs:

Aluno → com atributos nome e nota

Professor → com atributos nome e disciplina

Implemente Exibir() em cada struct.

No main, crie um slice de Exibivel contendo alunos e professores e percorra o slice imprimindo os resultados de Exibir().*/
