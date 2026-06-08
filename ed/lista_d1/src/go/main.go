package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// começa sempre com struct
type Node struct {
	//V maiusculo pq é um atributo público
	//v minusculo pq é atributo privado
	Value int
	next  *Node //aponta pro ponteiro do nó seguinte
	prev  *Node //aponta pro ponteiro do nó anterior
	//nil <- [10] <-> [20] <-> [30] -> nil
}

//class LList

type LList struct {
	root *Node //a raiz é um nó tbm
	size int
}

//root = nó sentinela
//quando a lista está vazia, ela aponta para si mesma

func NewLList() *LList {
	root := &Node{} //root aponta pra struct node(pra ela mesma no início)
	root.next = root
	root.prev = root

	return &LList{
		root: root,
		size: 0,
	}
}

// pro teste 1, preciso da func String
func (ll *LList) String() string {
	s := "["
	//inicia no primeiro nó válido da lista
	node := ll.root.next
	//evitar colocar a , antes do primeiro elemento
	//primeiro := true
	//percorrer a lista até a root(nó sentinela), que não tem um valor na lista
	for node != ll.root {
		//não sendo o primeiro elemento
		if node != ll.root.next {
			s += ", "
		}
		//adicionar o valor atual do nó na string
		s += fmt.Sprintf("%d", node.Value)
		//imprimir primeiro elemento c vírgula depois
		//primeiro = false
		//imprimir próximo nó
		node = node.next
	}

	s += "]"
	return s
}

//s := "["
// for i := 0; i < ms.size; i++ {
// 	if i > 0 {
// 		s += ", "
// 	}

// 	s += fmt.Sprintf("%d", ms.data[i])
// }

// s += "]"
// return s

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
