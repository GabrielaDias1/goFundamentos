package main

import "fmt"

func main() {
	nomes := []string{}

	nomes = append(nomes, "ana", "bia", "carlos") //append sabe inserir sozinho
	fmt.Println("Conteúdo do slice:", nomes)
	fmt.Println("tamanho slice criado", len(nomes))
	fmt.Println("capacidade slice criado", cap(nomes))

	//teste para ver se ao inserir outro ele se adapta e como se comporta
	nomes = append(nomes, "eduardo")
	fmt.Println("Conteúdo do slice:", nomes)
	fmt.Println("teste aumento tamanho slice criado", len(nomes))
	fmt.Println("teste aumento capacidade slice criado", cap(nomes))
    /*o Go aumenta a capacidade automaticamente para evitar precisar realocar 
	 memória a toda hora — ele faz isso crescendo a capacidade em "pedaços" 
	maiores para ser mais eficiente.*/
}

/*Adicionando elementos
Crie um slice vazio de strings e adicione nomes de amigos usando append.
Imprima a cada inserção o slice, seu tamanho e capacidade.*/
