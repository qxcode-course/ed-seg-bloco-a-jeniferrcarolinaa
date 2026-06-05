package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// para o primeiro teste, fazer o insert
// preciso da struct Set primeiro
type Set struct {
	data     []int
	size     int
	capacity int
}

//ah, e da func NewSet

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

//fun Sting

func (v *Set) String() string {
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
func (v *Set) Reserve(newCapacity int) {
	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}

//agora sim, func insert

func (v *Set) insert(index int, value int) error {
	//inserir um valor no index e deslocar o resto
	//verificar
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}
	//e se a capacidade for 0?
	//aumentar
	if v.capacity == 0 {
		v.Reserve(1)
	} else {
		v.Reserve(v.capacity * 2)
	}

	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]

	}
	v.data[index] = value
	v.size++
	return nil

}

func (v *Set) Insert(value int) {
	//adicionar elementos sem valores repetidos e na ordem crescente
	//assumindo que o valor será adicionado ao final do vetor
	//verificar se tem elementos repetidos

	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return
		}
	}

	posicao := v.size
	//aplicar um for para sair uma lista crescente
	for i := 0; i < v.size; i++ {
		if value < v.data[i] { //se nenhum valor for menor que o informado, vai pro final
			posicao = i
			break
		}
	}
	//chamar o insert interno
	_ = v.insert(posicao, value)
}

// func contains
func (v *Set) Contains(value int) bool {
	var achou bool
	achou = true
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return achou
		}

	}

	return !achou
}

// func erase
func (v *Set) Erase(value int) error {
	var index int
	index = -1

	//fazer verificação se o elemento existe
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			index = i
			break
		}
	}

	if index == -1 { //se não existe, publicar erro
		return fmt.Errorf("value not found")
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}

	v.size--
	return nil

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			err := v.Erase(value)
			if err != nil {
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
