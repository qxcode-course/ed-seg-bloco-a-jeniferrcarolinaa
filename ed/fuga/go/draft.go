// em Go, o loop é feito apenas com o for
// estrutura normal:
// for i:=0; i<5;i++{
// linha de código
// }
// estrutura estilo loop até um break (usando condições):
// for{
// condição
// break
// }
// estrutura estilo while:
// i:=0
// for i < 5{
// condição
// i++
// }
package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)

	for {
		if P == F {
			fmt.Printf("N\n")
			break
		}

		if H == F {
			fmt.Printf("S\n")
			break
		}
		//1 é antihorário
		//-1 é horário
		if D == 1 {
			F = F + 1
			if F == 16 {
				F = 0
			}
		}

		if D == -1 {
			F = F - 1
			if F == -1 {
				F = 15
			}
		}
	}
}
