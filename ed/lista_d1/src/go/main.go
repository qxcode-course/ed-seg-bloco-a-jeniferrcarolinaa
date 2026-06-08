package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
}

//root = nó sentinela
//quando a lista está vazia, ela aponta para si mesma

func NewLList() *LList {
	root := &Node{} //root aponta pra struct node(pra ela mesma no início)
	root.next = root
	root.prev = root

	return &LList{root: root}

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

// PushBack(value int)  // adiciona um novo nó com esse valor no fim da lista
func (ll *LList) PushBack(num int) {
	//preciso criar um novo nó (aquele a ser adicionado pelo usuário)E guardar o seu valor
	novo := &Node{Value: num}

	//último nó da lista vazia é o próprio root
	ultimo := ll.root.prev
	// 10,20(20 é o último)
	//adiciona nó 30
	//10,20,30 ->nil(ll.root)
	//o nó anterior ao novo (30) é o último anterior(20)
	novo.prev = ultimo

	//o nó próximo ao novo(que está no final da lista) é o root
	novo.next = ll.root

	//o antigo último(20) tem como next o novo(30)
	ultimo.next = novo

	//o novo elemento é o último da lista
	ll.root.prev = novo

}

// preciso do size pra correr o teste 3
func (ll *LList) Size() int {
	count := 0
	for n := ll.root.next; n != ll.root; n = n.next {
		count++
	}
	return count
}

// pushfront
// PushFront(value int) // adiciona um novo nó com esse valor no início da lista
func (ll *LList) PushFront(num int) { //empurra pra direita
	novo := &Node{
		Value: num,
	}

	//primeiro nó da lista vazia é o root
	primeiro := ll.root.next

	//quero adicionar um novo elemento
	//10,20
	//novo:30
	//30,10,20
	//anterior de 30(novo) é o root

	novo.prev = ll.root

	//próximo de 30 (novo) é o 20 (antigo primeiro)
	novo.next = primeiro

	//o próximo do root é o novo
	ll.root.next = novo

	//anterior do 20(antigo primeiro) é o novo(30)
	primeiro.prev = novo

}

// clear
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

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
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
