package main

import "fmt"

func imprimir(pessoas []int, pos int) {
	//imprime as filas das rodadas indicando quem tem a espada
	fmt.Printf("[ ")

	// preencher o vetor de pessoas
	for i := 0; i < len(pessoas); i++ {
		pessoas[i] = i + 1 // pessoa 1 no i 0, pessoa 2 no i 1, etc
		//quem está com a espada?
		if pos == i { //se a posição da pessoa for igual ao índice, ela tem a espada
			fmt.Printf("%d>", pessoas[i])
		} else { //caso n tenha a espada
			fmt.Printf("%d", pessoas[i])
		}

		//adicionar espaço entre os números

		if i < len(pessoas) {
			fmt.Print(" ")
		}
	}
	fmt.Printf("]\n")
}

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

	//enquanto houver mais de uma pessoa viva, o processo continua
	for len(pessoas) > 1 {
		imprimir(pessoas, pos)
		morto := (pos + 1) % len(pessoas)                       // pessoas restantes
		pessoas = append(pessoas[:morto], pessoas[morto+1:]...) //append junta os dois slices
		pos = morto % len(pessoas)
	}

	imprimir(pessoas, pos)

	//quem morre é o elemento na posição depois da pessoa com a espada
	//morto := pos + 1
	//sabendo que a fila é circular, se o morto = N, volta pro início
	//morto := (pos + 1) % len(pessoas) // pessoas restantes
	//retirar o morto da lista de pessoas
	//manter tudo antes do morto e juntar com tudo depois do morto
	//pegar o slice de pessoas antes do morto
	//aMorto := pessoas[:morto]
	//pegar o slice de pessoas depois do morto
	//dMorto := pessoas[morto+1:]
	//juntar os dois slices
	//pessoas = append(aMorto, dMorto...)

	//o ... é pra desagrupar o slice dMorto, adicionando elemento por elemento
	//para quem é transferida a espada?
	//para a pessoa que agora ocupa a posição que era do morto
	//pos = morto % len(pessoas)
	//precisa criar uma função para mostrar todo o processo, de cada rodada
}
