package main

import "fmt"

func main() {
	//macho +
	//fêmea -
	//entrada: número N de animais
	var N int
	fmt.Scan(&N)

	//guardar animais descasados em uma lista

	var nao_casados [50]int
	tam := 0
	pares := 0

	//ler o valor para cada animal no zoo
	for i := 0; i < N; i++ {
		var animal int
		fmt.Scan(&animal)
		achou := false
		//percorrer a lista de descasados
		for j := 0; j < tam; j++ {
			if nao_casados[j] == -animal {
				//marcar a posição do par no vetor nao_casados com 0
				nao_casados[j] = 0
				pares++ //p aumentar a quantidade de animais casados
				achou = true
				break
			}

		}
		//se n achou, adicionar na lista de não casados
		if !achou {
			nao_casados[tam] = animal
			tam++
		}
	}

	fmt.Println(pares)

}
