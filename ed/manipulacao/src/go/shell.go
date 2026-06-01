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
	_ = vet
	return nil
}

func reverse(vet []int) []int {
	_ = vet
	return nil
}

func unique(vet []int) []int {
	_ = vet
	return nil
}

func repeated(vet []int) []int {
	_ = vet
	return nil
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
