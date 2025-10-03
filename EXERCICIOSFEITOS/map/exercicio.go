package main

import "fmt"

func main() {
	agenda := make(map[string]string)
	agenda["gaby"] = "11959595959"
	agenda["eliana"] = "979797979"
	agenda["luana"] = "11926262626"

	fmt.Println("Contato 1 - Gaby:", agenda["gaby"])
	fmt.Println("Contato 2 - Eliana:", agenda["eliana"])
	fmt.Println("Contato 3 - Luana:", agenda["luana"])

	agenda["eliana"] = "970707070"
	fmt.Println("Mudança Contato 2 - Eliana:", agenda["eliana"])
}

/*agenda de contatos
Objetivo: praticar criação e acesso.
Tarefa: crie um map[string]string que associa nome → telefone.

Adicione 3 contatos.

Imprima o telefone de um nome específico.

Atualize o telefone de um contato e mostre antes/depois.*/
/*FORMULA:

variavel := make(map[TipoChave]TipoValor)*/
