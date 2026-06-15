package main

import (
	"fmt"
)

// pra entender o problema, vamos por partes, é confuso.
// temos uma string com dígitos e pontos
// tbm temos um limite maximo pra essa string
// temos que preencher esses pontos, de 0 a L
// ent se temos: 01.3. e L=3, temos:
// para preencher = 0,1,2,3
// o teste diz: 01320, 0 na posiçao 0 e 4. a distânia entre 0 e 4 é 4; 4>3, ent está correto
// que d>L
// como L = 3, quando vem um ., deve-se atender a regra de pos-3 até pos+3
// a pos que tem o . é a 2, logo, de -1 a 5
// a entrada deve ser em string, mas quero modificar o '.' para números, então entra rune tbm
// para fragmentar a string em partes e fazer a conversão
func Distancia(seq []rune, pos int, L int) bool {
	//variaveis até agr:
	//seq []rune, //array
	//pos int, // posição/indice
	//L int // limite
	//fazer algumas verificações, que retornam true ou false
	//ao chegar no fim, já foi tudo preenchido
	if pos == len(seq) {
		return true
	}

	//se a posição atual não for um ponto, segue
	//tem que seguir
	if seq[pos] != '.' {
		return Distancia(seq, pos+1, L) // continua seguindo
	}

	//correr todos os dígitos possíveis
	for d := 0; d <= L; d++ {
		//verificar se é válido
		//transforma 3 em '3', por exemplo
		digito := rune('0' + d)
		//pode colocar esse dígito no no lugar do ponto?
		//como colocar d na pos?
		//usando as regras já descritas no início em relação à distância
		valido := true //começa supondo que pode inserir
		//de pos-L a pos +L
		//não pode ter negativo
		//garantindo que n existe outro muito perto
		for i := pos - L; i <= pos+L; i++ {
			//se saiu dos limites da string, ignora o valor
			if i < 0 || i >= len(seq) {
				continue
			}

			//se encontrou o mesmo dígito, não pode
			if seq[i] == digito {
				valido = false
				break
			}
		}

		//se for válido, insere
		if valido {
			//coloca o digito na pos
			seq[pos] = digito
			//prossegue

			if Distancia(seq, pos+1, L) {
				return true
			}
		}
		//não funcionando, desfaz a tentativa
		seq[pos] = '.'
	}

	return false

}

func main() {
	//ler string
	var str string
	fmt.Scan(&str)
	//ler l
	var L int
	fmt.Scan(&L)
	//converet sting para rune
	seq := []rune(str)
	//chamar distanc
	Distancia(seq, 0, L)
	//imprimir string(seq)
	fmt.Println(string(seq)) //converte de volta para string
}
