package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Pair struct {
// 	One int//começa com letra maiuscula para ser usada fora do pacote
// 	Two int
// }

type Pair struct {
	One int
	Two int
}

// func occurr(vet []int) []Pair { //quantas vezes o nível de stress aparece // deve contar em valor absoluto
// 	// fazer um map para (1,2) -> nível,quantidade
// 	ocorrencia := make(map[int]int)
// 	_ = vet
// 	return nil
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func occurr(vet []int) []Pair { //criando uma função (func) que se chama occurr//
	//recebe um parâmetro chamado vet do tipo slice de int
	//precisa devolver um valor do tipo slice (lista) de Pair
	//precisa de abs, pra contar o numero que aquele nivel aparece independente se é homem ou mulher

}
func teams(vet []int) []Pair { //o nivel x tem y pessoas
	//(5,2) -> nivel 5 tem 2 pessoas
	_ = vet
	return nil
}

func mnext(vet []int) []int {
	//se tem mulher como vizinha, retorna 1, senão, retorna 0
	//se a posição for de mulher, retorna 0
	_ = vet
	return nil
}

func alone(vet []int) []int {
	//se tem vizinha, retorna 0, senão, retorna 1
	//se a posição for de mulher, retorna 0
	_ = vet
	return nil
}

func couple(vet []int) int {
	//homem e mulher com o mesmo nível de stress são casal
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	//tem que retorna o indice(pos) onde começa a sequencia procurada
	//se n houver, retorna -1
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	//remover pelo indice com erase[] elemento a ser removido
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	//remove pelo valor com clear[] valor a ser removido
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
