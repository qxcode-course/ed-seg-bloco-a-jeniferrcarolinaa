package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//mais teoria
//FIFO = First In, First Out = o primeiro que entra é o primeiro que sai
//head é elemento 1 e tail é o último elemento
//pop (dequeue) remove head
//push(enqueue) insere elemento após tail

func (q *Queue[T]) Enqueue(value T) {
	//estamos falando de nós, então temos que criar um primeiro
	novo := &Node[T]{Value: value}
	//se a fila estiver vazia, o nó 1 será a head e a tail ao mesmo tempo
	//se já houver elementos (q.head!=nil), o inserido(novo) aponta para o novo antigo(tail)

	if q.head == nil {
		q.head = novo
		q.tail = novo
	} else {
		q.tail.next = novo
		q.tail = novo
	}

	q.size++
}

func (q *Queue[T]) Dequeue() (T, bool) { //true se removeu, false se a lista está vazia
	var t T
	if q.head == nil {
		return t, false
	}
	//retornar a lista e true
	value := q.head.Value
	q.head = q.head.next

	//se a fila ficou vazia(head = nil), n tem tail tbm
	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return value, true
	//remove head
	//1,2,3,4
	//1 é a head
	//só fazer o segundo elemento ser a head
	//head.next = head
}

// func (q *Queue[T]) Peek() (T, bool)
// func (q *Queue[T]) Size() int
// func (q *Queue[T]) IsEmpty() bool
// func (q *Queue[T]) Clear()

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

// peek
// Visualização do primeiro elemento Retorna o primeiro elemento sem removê-lo.
func (q *Queue[T]) Peek() (T, bool) {
	var t T
	//verificar se a lista está vazia
	if q.head == nil {
		//fmt.Println("fail: fila vazia")
		return t, false
	}

	return q.head.Value, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
