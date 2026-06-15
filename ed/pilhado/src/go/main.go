package main

import (
	"bufio"
	"fmt"
	"os"
)

//fazer uma pilha
//como descobrir o caminho entre I e F?

// certo, passo 1 é criar a struct das posições da matriz
type Pos struct {
	l int
	c int
}

// passo 2 é criar as duas pilhas
// uma vai armazenar o caminho percorrido até o nó atual e
// a outra vai armazenar os pontos que identificamos serem becos sem saída.
// sabendo que são matrizes
func Pilhas(matrizes [][]rune, inicio Pos) { //var matrizes com inicio na linha e coluna
	//precisa guardar posições na matrix
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()
	//preciso especificar/adicionar um inicio(topo I)
	caminho.Push(inicio)
	//enquanto ainda houver espaço, olha pro primeiro espaço à frente
	for !caminho.IsEmpty() {
		espAtual := caminho.Top()

		_ = espAtual
		_ = becos
		//agr, encontrar os espaços válidos(cima, baixo, esquerda, direita)
		//muito parecido com a questão de queimando
		//#: parede
		//. = já visitado
		//ir marcando
		if matrizes[espAtual.l][espAtual.c] == 'F' {
			matrizes[espAtual.l][espAtual.c] = '.'
			break
		}
		//se chegou ao fim, para (cumpriu o objetivo)
		if matrizes[espAtual.l][espAtual.c] != '#' {
			matrizes[espAtual.l][espAtual.c] = '.'
		}
		//procurar as posições válidas
		//cima:=Pos{espAtual.l - 1, espAtual.c}
		//poderia fazer dessa forma, mas teria de verificar um a um
		//usar o vetor de direções {l,c}
		direcoes := []Pos{
			//cima
			{-1, 0},
			//baixo
			{1, 0},
			//direita
			{0, 1},
			//esquerda
			{0, -1},
		}
		//fazer um for para verificar se são válidas
		//percorrer cada direção,  vamos testar a estrutura do for in range
		var prox *Pos = nil
		for _, d := range direcoes {
			nl := espAtual.l + d.l
			nc := espAtual.c + d.c
			//verificar se estão dentro dos limites, igual a em queimando
			if nl < 0 || nl >= len(matrizes) {
				continue
			}
			if nc < 0 || nc >= len(matrizes[0]) {
				continue
			}
			//pode-se avançar quando encontra "" ou F
			if matrizes[nl][nc] == ' ' || matrizes[nl][nc] == 'F' {
				//pegar o primeiro "vizinho válido"
				primeiro := Pos{nl, nc}
				prox = &primeiro
				break
			}
		}
		//se já encontrou, vai pra prox posi
		if prox != nil {
			caminho.Push(*prox)
		} else {
			//se n encontrou, é um beco
			becos.Push(espAtual) //adiciona em becos
			caminho.Pop()        //retira da pilha de caminho
		}
	}
	//retirar marcações do beco

	for !becos.IsEmpty() {
		p := becos.Pop()
		if matrizes[p.l][p.c] == '.' {
			matrizes[p.l][p.c] = ' '
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	matrizes := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		matrizes = append(matrizes, line)
	}
	var inicio Pos
	for l := range matrizes {
		for c := range matrizes[l] {
			if matrizes[l][c] == 'I' {
				inicio = Pos{l, c}
			}
		}
	}
	Pilhas(matrizes, inicio)
	showMatriz(matrizes)
}

func showMatriz(matr [][]rune) {
	for _, linha := range matr {
		fmt.Println(string(linha))
	}
}
