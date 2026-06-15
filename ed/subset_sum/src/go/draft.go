//teoria backtraking
//queremos achar números em um conjunto.
//quando somados, deve resultar no inteiro que buscamos
//retorna true se houver uma soma que resulte no número
//lembra conceito de arvore de decisões
//exploração de possibilidades
//entra função recursiva
//toda função precisa saber: "onde estou"/índice e "quanto já somei"
//ent a pergunta é: a partir de onde estou consigo chegar ao valor alvo? se sim, retorna true
//variáveis envolvidas seriam: vet []int, indiceAtual, somaAtual, valorAlvo
//há dois casos: valorAlvo = somaAtual ee cheguei ao final mas não encontrei o valor (indiceAtual = len(vet))
//responder à pergunta se pega ou não o próximo número e o adiciona à soma
//chamadas rercusivas: uma função que chama ela mesma
//exemplo

package main

import "fmt"

// func contar(n int){
//     fmt.Println(n)
//     if n == 0{
//         return
//     }

//     contar(n-1)
//     //se n = 4
//     //contar(3)
//     //contar(2)
//     //contar(1)
//     //contar(0)
// }

//assim, precisa de uma função que chame ela mesma, em busca do valorAlvo
//e ela vai se chamando até que a soma seja igual ao alvo

func Backtraking(vet []int, indiceAtual int, somaAtual int, valorAlvo int) bool {
	//pensando nos dois casos
	//se a soma for igual ao alvo
	//se chegou até o fim e não encontrou o valor procurado
	if somaAtual == valorAlvo {
		return true
	}

	// caso  quando soma atual ultrapassa o valor alvo
	if somaAtual > valorAlvo || indiceAtual >= len(vet) {
		return false
	}

	// if indiceAtual >= len(vet) {
	// 	return false
	// }

	//depois de verificar, eu tenho duas chamadas:
	// a primeira é decidir se eu pego o elemento.
	//a outra é se eu ignoro o elemento
	pega := Backtraking(
		vet,
		indiceAtual+1,              //avança
		somaAtual+vet[indiceAtual], //soma
		valorAlvo,
	)
	ignora := Backtraking(
		vet,
		indiceAtual+1,
		somaAtual, //ignora
		valorAlvo,
	)
	//vai retornar um dos dois caminhos
	return pega || ignora
}

//agora a main
func main() {
	//preciso ler n e k
	//criar o vetor de tamanho n
	//ler n números (for)
	//chamar a func backtraking
	var n, k int
	fmt.Scan(&n, &k)
	vet := make([]int, n) //vetor de tam n

	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
	}

	//resultado := Backtraking(vet, 0, 0, k)
	if Backtraking(vet, 0, 0, k) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	// fmt.Println(Backtraking(vet, 0, 0, k))
	// retorna se é true ou false
}
