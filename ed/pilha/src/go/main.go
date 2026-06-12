package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//sobre pilhas
//LIFO - last in, first out - último a entrar é o primeiro a sair
//no pop, se tenho 1,2,3, fica 1,2

type Stack[T any] struct {
	data []T //o slice todo já é a pilha
	//tem a base e o topo(high) -> s.data[len(s.data) - 1]
}

func NewStack[T any](cap int) *Stack[T] { // Cria uma nova pilha com a capacidade inicial especificada
	return &Stack[T]{
		//criar slice vazio
		data: make([]T, 0, cap),
		//len(s.data) = 0, cap a ser definido
	}
}

// push pra teste 1
func (s *Stack[T]) Push(value T) {
	// Adiciona um valor ao topo da pilha utilizando append
	s.data = append(s.data, value)
}

// size, top e clear teste 2
// Peek() (T, error)        // Retorna o valor do topo da pilha sem removê-lo.
func (s *Stack[T]) Peek() (T, error) {
	//verificar se a lista está vazia
	var t T
	if s.Size() == 0 {
		return t, fmt.Errorf("stack is empty")
	}
	//topo
	return s.data[len(s.data)-1], nil
}

// Clear()                  // Limpa a pilha, removendo todos os elementos.
func (s *Stack[T]) Clear() {
	//sabendo que em slices[:x] significa criar um novo, como elementos até antes do índice x
	s.data = s.data[:0] //pegue tudo antes de 0, que é o próprio 0

}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

// pop
// Pop() error              // Remove o valor do topo da pilha.
func (s *Stack[T]) Pop() error {
	//verificar se tá vazio
	if s.Size() == 0 {
		return fmt.Errorf("stack is empty")
	}
	//1,2,3
	///pop
	//1,2
	s.data = s.data[:len(s.data)-1] //pegar até o penúltimo índice só
	return nil
}

func (s *Stack[T]) String() string {
	output := ""
	for i := range cap(s.data) {
		if i != 0 {
			output += ", "
		}
		if i < len(s.data) {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)
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
			cap, _ := strconv.Atoi(parts[1])
			v = NewStack[int](cap)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v)
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
