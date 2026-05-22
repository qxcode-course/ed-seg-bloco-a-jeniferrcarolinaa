package main

import "fmt"

func main() {
	var N, E int
	fmt.Scan(&N, &E)
	//sabe-se que em um vetor, se temos:
	// 1 2 3 4 5
	// indice:
	// 0 1 2 3 4
	//assim, pessoa 1 está no índice 0; pessoa 2 no índice 1; e assim por diante
	//conluindo, a pessoa E está no ínidice E-1
	//posição (ÍNDICE) da pessoa E = E-1
	//pos:= E-1
	//precisamos de um vetor para guardar as pessoas
	var pessoas = make([]int, N) // este é o slice de pessoas (tipo int) com tamanho/capacidade N
	pos := E - 1

	fmt.Printf("[ ")

	// preencher o vetor de pessoas
	for i := 0; i < N; i++ {
		pessoas[i] = i + 1 // pessoa 1 no i 0, pessoa 2 no i 2, etc
		//quem está com a espada?
		if pos == i { //se a posição da pessoa for igual ao índice, ela tem a espada
			fmt.Printf("%d>", pessoas[i])
		} else { //caso n tenha a espada
			fmt.Printf("%d", pessoas[i])
		}

		//adicionar espaço entre os números

		if i < N {
			fmt.Print(" ")
		}
	}

	fmt.Printf("]\n")

}
