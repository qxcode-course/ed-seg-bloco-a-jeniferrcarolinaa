package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

//adiantar func String

func (ms *MultiSet) String() string {
	s := "["
	for i := 0; i < ms.size; i++ {
		if i > 0 {
			s += ", "
		}

		s += fmt.Sprintf("%d", ms.data[i])
	}

	s += "]"
	return s
}

//para aumentar a capacidade do vetor, agr é o expand

func (ms *MultiSet) expand(newCapacity int) {
	newData := make([]int, newCapacity)
	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}
	ms.data = newData
	ms.capacity = newCapacity
}

//func search

func (ms *MultiSet) search(value int) (bool, int) {
	//fazer o equivalente ao bettersearch
	low := 0
	high := ms.size - 1

	for low <= high {
		midlle := (low + high) / 2

		if ms.data[midlle] == value { //se encontrou o valor procurado
			//preciso resolver o problema do registro da última ocorrência do elemento x
			last := midlle //guarda a pos da última ocorrência
			//garantir que esteja dentro dos limites do vetor
			//&& verificar se o próximo elemento é igual ao valor procurado
			for last+1 < ms.size && ms.data[last+1] == value {
				last++ //incrementa a pos, foi achado um outro elemento igual
			}
			return true, last

		}

		if value < ms.data[midlle] {
			high = midlle - 1
		} else {
			low = midlle + 1
		}
	}

	return false, low
}

// insert
func (ms *MultiSet) insert(value int, index int) error {
	//verificação
	if index < 0 || index > ms.size {
		fmt.Errorf("index out of range")
	}

	//expandir capacity quando necessário
	if ms.capacity == 0 {
		ms.expand(1)
	} else {
		ms.expand(2 * ms.capacity)
	}

	//deslocar elementos pra direita
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++

	return nil
}

// agora o Insert
func (ms *MultiSet) Insert(value int) {
	achou, posicao := ms.search(value)
	//inserir após a última ocorrencia
	//função search retorna 2 valores, um bool e um inteiro
	if achou { //incrementar a repetição
		ms.insert(value, posicao+1) //value,index
	} else {
		ms.insert(value, posicao)
	}

}

// contains
func (ms *MultiSet) contains(value int) bool {
	var achou bool
	achou = true
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			return achou
		}
	}

	return !achou
}

// erase
func (ms *MultiSet) erase(index int) error {
	//verificação
	if index < 0 || index > ms.size {
		fmt.Errorf("index out of range")
	}

	//fazer o contrário de insert, deslocar os elementos pra esquerda
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
	return nil
}

//Erase

func (ms *MultiSet) Erase(value int) error {
	//ideia semelhante ao Insert
	achou, index := ms.search(value)

	if !achou {
		return fmt.Errorf("value not found") //caso não se encontre o valor procurado
	}
	//chamar o erase
	return ms.erase(index)

}

// count
// + Count(value: int): int                  ' Retorna o número de ocorrências do valor no multiconjunto
func (ms *MultiSet) count(value int) int {
	contador := 0 //estabelecer um contador
	for i := 0; i < ms.size; i++ {
		if ms.data[i] == value {
			contador++
		}
	}

	return contador
}

// unique
// Unique(): int                           ' Retorna o número de valores distintos no multiconjunto
func (ms *MultiSet) unique() int {
	//criar contador e ir increme. por elemento diferente
	//se n tiver elementos
	if ms.size == 0 {
		return 0
	}
	contador := 1 //tem pelo menos um elemento distinto
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			contador++
		}
	}

	return contador
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.count(value))
		case "unique":
			fmt.Println(ms.unique())
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
