package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//relembrando o binarysearch da questão anterior
//realiza busca binária e retorna o índice do valor, ou-1 se não for encontrado

func BetterSearch(slice []int, value int) (bool, int) {
	//high é o final do vetor
	//low é o índice inicial, ou seja, 0
	//middle é o meio do vetor
	//a diferença entre encontrar o elemento pelo for e pelo binary search é que o custo de operação
	high := len(slice) - 1 //útimo índice
	low := 0               //

	//middle :=  (low +high)/ 2

	for low <= high { //enquanto ainda tiver elemento a ser procurado
		middle := (low + high) / 2
		if value == slice[middle] {
			return true, middle
		}

		if value < slice[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}

	return false, low //onde deveria estar
	//se o valor procurado estiver antes do meio do vetor
	// if value < slice[middle]{
	// 	high = middle -1
	// }
	//estando depois
	// if value > slice[middle]{

	// }

	//_, _ = slice, value

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
