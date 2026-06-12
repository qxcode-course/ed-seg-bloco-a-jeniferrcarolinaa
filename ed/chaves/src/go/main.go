//esta meio confuso mas vamos lá
//são 15 jogos
//16 times diferentes
//2 times (duas letras) jogam um jogo por vez
//vai de A a P(16 letras)
//A - B
//2   3 />representa a pontuação de cada um
//8 jogos início
//dps 4
//dps 2
//dps 1
//ganhador
// 2 3 - B
// 3 1 - C
// 2 1 E
// 1 5 H
// 2 4 J
// 5 4 K
// 5 2 M
// 2 5 P

// 1 4 times B e C - C
// 1 5 E e H - H
// 2 5 J e K - K
// 3 4 M e P - P

// 1 2 C e H - H
// 5 3 K e P - K

// 5 1 H e K - H ganha
package main

import (
	"fmt"
)

func main() {
	//primeiro eu tenho que criar uma fila com os 15 times
	//NewQueue é a função de criação de fila que vou usar
	//como são letras, o tipo é string
	times := NewQueue[string]()
	//começando em A, indo até P
	//pesquisar sobre como adicionar strings
	//UniCode , cada caractere representa um
	for caractere := 'A'; caractere <= 'P'; caractere++ {
		//adicionar na fila
		//times.push()
		//quero a string, não o unicode
		times.Enqueue(string(caractere))
	}
	//depois que todos os times (representados por letras) estiverem na fila, adicionamos os gols
	for i := 0; i < 15; i++ {

		//cada jogo pega duas letras por vez
		time1 := times.Dequeue()
		time2 := times.Dequeue()

		//gols associados a cada time
		var gol1, gol2 int
		fmt.Scan(&gol1, &gol2)

		//selecionar o maior gol, e este entra pra fila de times de novo
		//meio que vai cortando a quantidade pela metade até sobrar um, o vencedor
		if gol1 > gol2 {
			times.Enqueue(time1)
		} else {
			times.Enqueue(time2)
		}
	}

	//no final sobre um né
	ganhador := times.Dequeue()
	fmt.Println(ganhador)
}
