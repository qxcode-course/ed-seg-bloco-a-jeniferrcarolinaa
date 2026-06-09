package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	s := "[ "

	//começar de um nó válido
	for node := l.Front(); node != l.End(); node = node.next {
		//adicionar vírgula
		// if node != l.root.next {
		// 	s += ", "
		// }

		s += fmt.Sprintf("%d", node.Value)

		//sword está no node informado
		if node == sword {
			s += ">"
		}
		s += " "

	}

	s += "]"
	return s

}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.Next() == l.End() { //quando chega no final
		return l.Front() //volta ao primeiro
	}
	return it.Next() //tem que ir passando pro próximo
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
