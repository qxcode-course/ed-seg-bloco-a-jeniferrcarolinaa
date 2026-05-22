package main

import "fmt"

func main() {
	var Q int
	var D string
	fmt.Scan(&Q, &D)
	//ler posições x e y
	var x [100]int
	var y [100]int
	for i := 0; i < Q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	//cada gomo subsequente é a posição anterior + um passo na direção D
	for i := Q - 1; i > 0; i-- {
		x[i] = x[i-1]
		y[i] = y[i-1]
	}
	//a depender da direção D, calcular a nova posição da cabeça

	switch D {
	case "L":
		x[0]--
	case "R":
		x[0]++
		//eixo y aumenta para baixo
	case "U":
		y[0]--
	case "D":
		y[0]++
	}

	for i := 0; i < Q; i++ {
		fmt.Println(x[i], y[i])
	}

}
