package main

import "fmt"

func main() {
	//ler a quantidade de figurinhas do album
	var qAlbum int
	var qTotal int
	fmt.Scan(&qAlbum)
	fmt.Scan(&qTotal)
	//em ordem crescente, ler as figurinhas do album
	//usar o for para ler as figurinhas do album
	var contador [51]int //vetor que conta a quantidade repetida da figu
	for i := 0; i < qTotal; i++ {
		var fig int    //figurinha que está sendo lida
		fmt.Scan(&fig) //guarda a figurinha lida
		//saber quantas vezes a figurinha foi repetida
		contador[fig]++
	}

	//imprimir as figurinhas repetidas
	achou := false //ainda não achou figurinhas repetidas
	//percorrer figurinhas do álbum
	for i := 1; i < qTotal; i++ {
		for j := 1; j < contador[i]; j++ {
			if achou {
				fmt.Print(" ") //
			}
			fmt.Print(i)
			achou = true //achou figurinhas repetidas
		}
	}

	fmt.Println() //quebra de linha

	if achou == false {
		fmt.Print("N")
	}

	achou = false
	for i := 1; i <= qAlbum; i++ {
		if contador[i] == 0 { // quando a fig não foi encontrada
			if achou {
				fmt.Print(" ")
			}
			fmt.Print(i)
			achou = true

		}

	}

	fmt.Println()

	if achou == false {
		fmt.Print("N")
	}
}
