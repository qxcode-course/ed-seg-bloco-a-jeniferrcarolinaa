package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	//mesmo que antes, mas adiciona o root
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
} //NÃO SEI COMO FAZ PRA RODAR AQUIII TU FAZ UM COMMIT PRA CADA VEZ QUE TENTA RODAR O CÓDIGO?

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root}
}

// implementar Next() e Prev()
func (n *Node) Next() *Node {
	//localizar o prox root
	//verificar se próximo nó é o sentinela
	if n.next == n.root {
		return nil
	}
	//não sendo, parte para o próximo
	return n.next
}

func (n *Node) Prev() *Node {
	//mesma lógica
	if n.prev == n.root {
		return nil
	}

	return n.prev
}

//importante "proteger" o ll.root.next e o "".prev

func (ll *LList) Front() *Node {
	//se a lista estiver vazia, size é 0
	if ll.root.next == ll.root {
		return nil
	}

	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.root.prev == ll.root { //se o anterior for root
		return nil //nulo
	}
	return ll.root.prev
}

//fazer a func string

func (ll *LList) string() string {
	s := "["
	//começa do primeiro nó válido
	//vai até node não ser o root(nil)
	//incrementa para o próximo
	for node := ll.Front(); node != nil; node = node.Next() {
		//adicionar vírgula após o primeiro node
		if node != ll.Front() {
			s += ", "
		}

		//adicionar os nós às listas
		s += fmt.Sprintf("%d", node.Value)
	}

	s += "]"
	return s
}

// func size
func (ll *LList) Size() int {
	count := 0
	for n := ll.root.next; n != ll.root; n = n.next {
		count++
	}

	return count

}

// func pushfront
func (ll *LList) PushFront(num int) {
	novo := &Node{Value: num, root: ll.root}
	primeiro := ll.root.next
	novo.prev = ll.root
	novo.next = primeiro
	ll.root.next = novo
	primeiro.prev = novo
	ll.size++
}

// pushback
func (ll *LList) PushBack(num int) {
	//mesmo que D1, basicamente
	novo := &Node{Value: num, root: ll.root}
	ultimo := ll.root.prev
	novo.prev = ultimo
	novo.next = ll.root
	ultimo.next = novo
	ll.root.prev = novo
	ll.size++
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

// implementar o search
// modelo iterador
func (ll *LList) Search(value int) *Node {
	//começa pelo primeiro valor válido do nó
	for node := ll.Front(); node != nil; node = node.Next() {
		//verificar se encontrou o valor
		if node.Value == value {
			return node //retorna o nó
		}
	}

	return nil
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
			fmt.Println(ll.string())
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
