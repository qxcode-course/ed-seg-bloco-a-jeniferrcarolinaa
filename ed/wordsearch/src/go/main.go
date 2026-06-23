//estudo da questão, um segundo
//backtracking semelhante a labirinto
//acessando a matriz, acesso a palavra
//onde estou na matriz -> qual a letra (que forma a palavra) estou procurando?
//estando em (l,c), consigo formar a palavra a partir da letra(x)?
//caso 1: se achei a palavra inteira: index = len(word)
//caso 2: posição inválida (se saiu da matrix, return false né)
//caso 3: letra não bate com a procurada grid[][]!=word[index], retornando false
//resumidamente: se a letra bate, marcar como visitada. tentar as 4 direções. 
// //nenhuma funcionando, desfaz
//importante: como funciona uma matriz de letras? há algo de especial? 
//estudar esse tópico
//n existe tipo char em go, mas tem aquela tabela de asc, algo assim
//o mais prox é rune e byte, byte é mais usado
//estrutura geral: grid [][]byte, ou seja
//grid [0][0] = 'A'
//grid [0][1] = 'B'
//etc
//grid [linha][coluna]byte
//ent
//palavra (string) = 'ABCD' acessa palavra[0] = 'A', palavra[1] = 'B' (?) estudiar
//passo 1 é criar func auxiliar que faz a procura profunda de letras que formam a palavra,
//com grid, l, c, index, palavra, como parâmetros

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
//essa aqui procura possíveis começos, possíveis pontos de começo, e chama a aux p ver se é possível formar a palavra com a letra atual (index)
//percorre a matriz -> encontra o ponto inicial -> chama a func aux

//func auxiliar
func aux(grid [][]byte, word string, l, c, index int) bool{//faltou uma virg, consertei 
	//caso 1:
	//se achou a palavra toda já, retorna true
	if index == len(word){
		return true
	}

	//caso 2, saiu além dos limites da matrix, retornando falso
	//labirinto, dar uma espiada pra lembrar
	if l<0||l>=len(grid) || c<0|| c>=len(grid[0]){
		return false
	}

	//caso 3, quando a letra que esta lendo agr não está dentro das procuradas pra formar a palavra alvo
	if grid[l][c] != word[index]{
		return false
	}

	//marcar as pos l e c já visitadas
	//pescando das outras questões um pouco, one sec
	//temos que garantir que a letra já 'visitada' não seja novamente
	//guardar a atual em uma var temp, marcando c um símbolo qualquer
	//se a letra for != deste símbolo, quer dizer que n visitou ainda ela
	//ok
	//depois se tenta os vizinhos (cima, baixo, esq e dir)

	//guarda a letra lida de agora
	temp := grid[l][c]
	//símbolo
	grid[l][c] = '#'
	//tenta as 4 direcoes
	//igual a labirinto e arvores
	//aux(grid [][]byte, word string, l, c index int)
	encontrando:=
		aux(grid, word, l-1, c, index + 1)||//cima
		aux(grid, word, l+1, c, index + 1)||//baixo
		aux(grid, word, l, c+1, index + 1)||//esq
		aux(grid, word, l, c-1, index + 1)//dir
	//dps de ler, devolve a temp
	grid[l][c] = temp
	return encontrando
	//yay, passamos nos teste 1 risos
}
func exist(grid [][]byte, word string) bool {
	//percorrer a matriz de letras pra achar o ponto inicial e depois chama a aux
	//como percorrer uma matriz de letras?
	//n tem muita dif, só percorrer linhas primeiro, e depois as coluna
	for l:=0; l<len(grid);l++{
		for c:=0; c<len(grid[0]);c++{
			//se a posição atual eh igual à primeira da palavra alvo
			//word[0]
			if grid[l][c] == word[0]{
				//recursao, duh
				if aux(grid, word, l, c, 0){
					return true
				}
			}
		}
		//return false
		//oops, falta um teste, da word SEE
		//problema é que tem mais de um início (dois Ss?)
		//vacilo, n vi a identação ohi
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
