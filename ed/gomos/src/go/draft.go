package main

import "fmt"

func main() {
	var Q int
	var D string
	fmt.Scan(&Q, &D)
	//ler posições x e y
	var x, y int
	for i := 0; i < Q; i++ {
		fmt.Scan(&x, &y)
		//a depender da direção D, calcular a nova posição
		switch D {
		case "L":
			x--
			if Q%2 != 0 && Q > 1 {
				y--
				x++
			}
		case "R":
			x++
		case "U":
			y++
		case "D":
			y--
		}

		fmt.Println(x, y)
	}
}
