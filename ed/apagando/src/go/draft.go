package main

import "fmt"

func main() {
	//quantidade de pessoas
	var N int
	fmt.Scan(&N)
	//identificador de cada um
	var entraram []int
	//cada pessoa tem um id
	//cada N, um id armazenado em uma lista
	for i := 0; i < N; i++ {
		var id int
		fmt.Scan(&id)
		entraram = append(entraram, id) // para adicionar cada pessoa na lista de ids
	}

	//M, numero de ids a serem retirados da lista de ids
	var M int
	fmt.Scan(&M)

	//var sairam []int
	// o id não pode ser retirado mais de uma vez
	// como garantir isso? map para verificar se o id já foi retirado
	//saiu[idR] = true
	//id saiu[idR] = true(
	sairam := make(map[int]bool) // crie um map com o tipo (de chave) int e (valor) bool
	//ids a serem retirados
	for i := 0; i < M; i++ {
		var idR int
		fmt.Scan(&idR)
		sairam[idR] = true // adicionar cada id retirado na lista de ids que sairam
	}

	//imprimir os ids que ficaram

	var ficaram []int
	for i := 0; i < N; i++ {
		//id da lista que entraram
		id := entraram[i]
		//verificar se id saiu ou não
		if !sairam[id] {
			ficaram = append(ficaram, id)
		}
	}
	//imprimir lista de ficaram
	for i := 0; i < len(ficaram); i++ {
		fmt.Printf("%d ", ficaram[i])
	}

}
