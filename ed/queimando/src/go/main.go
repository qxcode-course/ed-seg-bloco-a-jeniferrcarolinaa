package main

import (
	"bufio"
	"fmt"
	"os"
)

//revisar matriz
//DFS: busca em profundidade
// >>>>>>>> INSERT
// 2 3 1 1
// #.# começou no ponto e se expandiu para os lados
// .##
// ======== EXPECT
// #.o
// .oo
//o fogo começa na arv atual, pega as vizinhas, que pega as outras vizinhas
//# --> o

//em burnTrees, temos as pos l e c, criamos ent uma struct com elas(representam onde começa o fogo)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]() //struct se chama Pos
	//_ , _ , _ = mat, l, c
	//criar a pilha e adicionar as posições nela
	stack.Push(Pos{l: l, c: c})

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:(isEmpty)
	for !stack.IsEmpty() {
		pos := stack.Pop()
		//fmt.Println(pos)
		//grid[linha][coluna]
		if pos.l < 0 || pos.l >= len(grid) { //n pode passar a matriz(rune), ent deve ignorar tbm
			continue //ignora
		}
		if pos.c < 0 || pos.c >= len(grid[0]) {
			continue
		}

		//precisa saber se é arvore ou n, e depois transforma em o se tiver de queimar
		if grid[pos.l][pos.c] != '#' {
			continue //é um . ou já está queimada
		}
		grid[pos.l][pos.c] = 'o' //se não for, queima
		//e se queimou, pega as vizinhas
		//cima, linha de cima, mesma coluna
		stack.Push(Pos{
			l: pos.l - 1, c: pos.c,
		})
		//baixo, linha de baixo, mesma coluna
		stack.Push(Pos{
			l: pos.l + 1, c: pos.c,
		})
		//direita mesma linha, prox coluna
		stack.Push(Pos{
			l: pos.l, c: pos.c + 1,
		})
		//esquerda mesma linha, coluna ante
		stack.Push(Pos{
			l: pos.l, c: pos.c - 1,
		})
	}
	//   - retirar o elemento do topo(pop)
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha
	//o problema é que os vizinhos da arvore na pos(0,0) não existem d ecerta forma(pos -1)
	// e se for na outra borda, não pode ultrapassar o valor max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
