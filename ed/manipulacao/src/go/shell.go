package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
func getMen(vet []int) []int {
	//criar um vetor para armazenar homens
	var men []int
	//ler os elementos do vetor de entrada e percorrer cada um deles
	for i := 0; i < len(vet); i++ {
		//verficar se o elemento é positivo (homem)
		if vet[i] > 0 {
			//adicionar a lista de homens
			men = append(men, vet[i])
		}
	}

	return men
}

func getCalmWomen(vet []int) []int {
	// criar um vetor pra armazenar calm women
	var calmWomen []int
	//ler vetor de entrada
	for i := 0; i < len(vet); i++ {
		//o elemento é negativo?
		//nivel de stress < 10?
		//vou ter que fazer uma função abs pra ver o nivel absoluto de stress
		if vet[i] < 0 {
			if abs(vet[i]) < 10 {
				calmWomen = append(calmWomen, vet[i])
			}
		}
	}

	return calmWomen
}

func sortVet(vet []int) []int { // uma lista de inteiros devolve  a mesma lista de inteiros, mas ordenada
	//imprimir a lista dada em ordem crescente
	//fazer bubble sort

	for i := 0; i < len(vet); i++ {
		//comparar o atual com o proximo
		for j := i + 1; j < len(vet); j++ {
			//se o atual eh maior que o proximo, trocar de lugar
			if vet[i] > vet[j] {
				//guardar atual em uma var temporaria
				temporaria := vet[i]
				//colocar o proximo no lugar do atual
				vet[i] = vet[j]
				vet[j] = temporaria
			}
		}
	}

	return vet

}

func sortStress(vet []int) []int {
	//ordenar a lista de acordo com nivel d stress, usar a funcao abs pra comparar cada elemento (sinal n importa)
	for i := 0; i < len(vet); i++ {
		for j := i + 1; j < len(vet); j++ {
			if abs(vet[i]) > abs(vet[j]) {
				//fazer a mesma coisa do bubble
				temporaria := vet[i]
				vet[i] = vet[j]
				vet[j] = temporaria
			}
		}
	}

	return vet
}

func reverse(vet []int) []int { //do maior para o menor
	//bublle sort invertido
	//o usuário dá a lista a ser invertida
	//imprimo a mesma lista, e depois imprimo novamente, só que do maior pro menor
	//ordenar a lista do maior pro menor
	// for i := 0; i < len(vet); i++ {
	// 	for j := i + 1; j < len(vet); j++ {
	// 		//imprimir a lista primeiro
	// 		if vet[i] < vet[j] {
	// 			temporaria := vet[i]
	// 			vet[i] = vet[j]
	// 			vet[j] = temporaria

	// 		}
	// 	}
	// }

	//return vet

	//tem que começar do ultimo elemento ao primeiro, é para inverter a ordem, não ordenar
	var invertido []int
	for i := len(vet) - 1; i >= 0; i-- {
		invertido = append(invertido, vet[i])
	}

	return invertido
}

func unique(vet []int) []int {
	//retornar uma nova lista sem repetição de valores
	//colocar oselementos que se repetem em um map
	vistos := make(map[int]bool)
	var unicos []int
	//para cada elemento, verficar se já existe (no map)
	for i := 0; i < len(vet); i++ {
		//se não apareceu ainda, colocar na lista nova e depois indicar como visto
		//sabendo que o map de repetidos tem inteiros que retornam true ou false (booleano):
		if !vistos[vet[i]] {
			unicos = append(unicos, vet[i])
			vistos[vet[i]] = true
		}

		//se já apareceu, ignora
	}

	return unicos

}

func repeated(vet []int) []int {
	//retornar uma nova lista com os elementos que se repetem
	//elemento x ja apareceu quatas vezes?
	//map de inteiro que retorna um inteiro representante da quantidade de vezes q x apareceu
	contagem := make(map[int]int)
	//percorrer a lista dada
	//incrementar a contagem de cada elemento sempre que ele aparece
	for i := 0; i < len(vet); i++ {
		contagem[vet[i]]++
	}

	var repetidos []int //lista de todas as repetições
	//map n tem ordem, ent tem que percorrer de novo o map
	// for i := 0; i < len(vet); i++ {
	// 	//se apareceu mais de uma vez, colocar em repetidos
	// 	if contagem[vet[i]] > 1 {
	// 		repetidos = append(repetidos, vet[i])
	// 	}
	// }

	//mudança de planos: tenho que percorrer o map, pegando a quantidade do valor que se repete e depois diminuir em uma unidade
	//estrtutura para percoreer o map:
	// for chave, valor := range contagem {

	for valor, quantidade := range contagem {
		if quantidade > 1 {
			//percorrer a quantidade de vezes que o valor se repete na lista
			for r := 0; r < quantidade-1; r++ {
				repetidos = append(repetidos, valor)
			}
		}
	}
	return repetidos
	//foda!!!, tem que ser repetidos -1, testes indicam

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
