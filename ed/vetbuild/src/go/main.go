package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	//função para criar um vetor
	return &Vector{
		data: make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		//é assim que se cria um vetor
		size:     0,
		capacity: capacity,
	}
}

// criar função status
// retorns o tamanho e a capacidade do vetor
func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

// criar função String pra imprimir o vetor
func (v *Vector) String() string { // função se chama string, recebe um ponteiroi Vector e devolve tipo string

	s := "["

	for i := 0; i < v.size; i++ {

		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%d", v.data[i])
	}

	s += "]"
	return s
}

// fun Reserve para aumentar a capacidade do vetor
func (v *Vector) Reserve(newCapacity int) {
	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}

// pushBack para adicionar um elemento ao final do vetor
func (v *Vector) PushBack(value int) {
	//o vetor tem capacidade pra adicionar o elemento?
	//se tiver, adicionar o elemento e aumenta o size
	//não tendo, aumentar capacity com reserve e depois adicionar o elemento
	//fazer verificação
	if v.size == v.capacity {
		//adicionar o elento no vetor v, na posição size
		//value é o elemento a ser adicionado
		v.Reserve(v.capacity * 2)
		//incrementar o size
	}
	//não tendo capacidade, aumentar a capacidade do vetor

	//agr posso adicionar o elemento
	v.data[v.size] = value
	v.size++
	//aff, vou ter que criar outra função pra aumentar a capacidade do vetor
}

func (v *Vector) Get(index int) int {
	return v.data[index] //elemento no indice/posição index
	//preciso criar a função At
}

func (v *Vector) At(index int) (int, error) {
	//func para verificar se o index é válido
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range") //0 é só pra compilar, tipo do return é o int
	}

	return v.Get(index), nil //o nil é a ausência de erro

}

// ainda falta o set pra passar no teste
func (v *Vector) Set(index int, value int) error {
	//fazer verificação
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	//definir o valor no index especificado
	v.data[index] = value
	//sem erro
	return nil
}

// func clear agr
func (v *Vector) Clear() {
	//limpar o vetor, é size 0
	v.size = 0
	//

}

// func PopBack, Remove e retorna o último elemento do vetor. Retorna um erro se o vetor estiver vazio
func (v *Vector) PopBack() error {
	//
	//remove e retorna o ultimo elemento
	//mensagem de erro se estiver vazio
	if v.size == 0 {
		return fmt.Errorf("vector is empty")
	}
	v.size--
	return nil
}

// func insert
func (v *Vector) Insert(index int, value int) error {
	//inserir um valor no index e deslocar o resto
	//verificar
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}

	if v.size == v.capacity {
		v.Reserve(v.capacity * 2)
	}

	//tem que copiar os elementos a serem deslocados para a posição à direita, para que n se perca
	//de trás pra frente, assumindo que i começa no final
	//para quando i atinge o index(indice do valor que quero inserir)
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}

	v.data[index] = value
	v.size++
	return nil

}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--
	return nil
}

func (v *Vector) IndexOf(value int) int {
	//retorna o índice do valor especificado
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i
		}
	}
	return -1

}

func (v *Vector) Contains(value int) bool {
	//percorrer o vetor
	//confirmar que o valor está no vetor
	var achou bool
	achou = true
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return achou
		}
	}

	return !achou
}

func (v *Vector) Slice(start int, end int) *Vector {
	//retorna um novo vetor
	novoVetor := &Vector{} //do ponteiro *Vector
	//começa em start, termina em end
	//sabendo que -x adiciona mais elementos
	if end < 0 {
		end = v.size + end
	}

	novoVetor.data = v.data[start:]  //corta tudo que vem antes do i start
	novoVetor.capacity = end - start // atualizar capacidade
	novoVetor.size = end - start     //atualizar o tamanho
	return novoVetor
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
