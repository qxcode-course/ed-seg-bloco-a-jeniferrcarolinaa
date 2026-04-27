package main

import "fmt"

func main() {
	//Entrada de N competidores
	var N int
	fmt.Scan(&N)
	//Entendendo o loop em Go:
	//for i:= 0; i<; i++

	indiceV := -1
	melhorP := 201

	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)

		if A >= 10 && B >= 10 {
			p := A - B
			if p < 0 {
				p = -p
			}
			//como fazer o módulo da diferença entre dois números ?
			//não existe pra inteiro, apenas para float64. seria:
			//p := math.Abs(float64(A - B))

			if p < melhorP {
				melhorP = p
				indiceV = i
			}
		}
	}

	if indiceV == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(indiceV)
	}

}
