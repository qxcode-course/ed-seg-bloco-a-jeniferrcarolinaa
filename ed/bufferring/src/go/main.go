package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data     []int
	front    int
	size     int
	capacity int
}

//estudar teoria né gente
//o front é o índice em que se encontra o primeiro valor
//size é a quantidade de elementos válidos
//para dar "a volta" : (index+capacity)%capacity = 5,6,7
//pra passar no primeiro teste, fazer primeiro o Len(), que retorna o tam do size

func (b *Deque) String() string {
	result := []string{}
	for i := range b.size {
		val := b.data[(b.front+i)%b.capacity]
		result = append(result, fmt.Sprint(val))
	}
	return "[" + strings.Join(result, ", ") + "]"
}

func (b *Deque) Debug() string {
	result := make([]string, b.capacity)
	for i, _ := range result {
		result[i] = " _"
		if i == b.front {
			result[i] = ">_"
		}
	}
	for i := range b.size {
		index := (b.front + i) % b.capacity
		val := b.data[index]
		prefix := " "
		if i == 0 {
			prefix = ">"
		}
		result[index] = fmt.Sprintf("%s%d", prefix, val)
	}
	return strings.Join(result, " |")
}

func (b *Deque) Len() int {
	//retorna o número de elementos
	return b.size
}

// pushback
// PushBack(v int)      // insere valor no fim
// precisa do resize
func (b *Deque) resize(newCap int) {
	//dobra a capacidade quando ela atinge o máximo
	//tenho que criar um novo vetor pq é um círculo que n entende onde é o início e o fim,
	//daí criar um for para copiar elemento por elemento no novo vetor, para especificar i e f
	newV := make([]int, newCap)
	for i := 0; i < b.size; i++ {
		//pegar os elementos do vetor antigo um por um, a patir do b.front + 0. depois +1, sempre avançamdo
		newV[i] = b.data[(b.front+i)%b.capacity]
	}

	//o antigo agora é o newV
	b.data = newV

	//cap antiga agr é newCap
	b.capacity = newCap

	//e o front passa a ser o 0
	b.front = 0
}
func (b *Deque) PushBack(value int) {
	//>10,20,30_
	//pushback 5
	//>10,20,30,5
	//fórmula do fim: (front + size) % capacity
	//pode acontecer de estar cheio (capacidade atingida), sendo necessário chamar o resize()
	//incrementar o size depois
	//index é a posição onde quero adicionar o elemento, neste caso, o fim
	//fazer a verificação e chamar o resize
	//se o tamanho for igual à capacidade, dobrar capacidade
	if b.size == b.capacity {
		b.resize(b.capacity * 2)
	}
	index := (b.front + b.size) % b.capacity
	//adicionar o valor no index do vetor
	b.data[index] = value
	b.size++
}

// fazer popfront para passar no segundo teste
// PopFront()           // remove valor do início
func (b *Deque) PopFront() error {
	//não desloca elementos
	//pra remover o primeiro, é só mudar o front e decrementar o size
	//tem mensagem de erro
	//verificar se o buffer está vazio
	if b.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}

	//front = 0
	//size = 4
	//1,2,3
	//popfront
	//front = 1
	//size = 3
	//result = 2,3

	//aumentar em uma unidade o front
	b.front = (b.front + 1) % b.capacity
	b.size--
	return nil
}

// legal, passaram dois de vez!
// agora é o clear() e o pushfront()
func (b *Deque) Clear() {
	//só zerar o size
	b.size = 0
	//quando fizer o pushback, queremos que o primeiro elmento seja adicionado no front 0
	b.front = 0
}

// PushFront(v int)     // insere valor no início
// lófica semelhante ao push back, mas c outra formula
func (b *Deque) PushFront(value int) {
	//fazer verificação da capacidade
	if b.capacity == 0 {
		b.resize(1)
	} else if b.size == b.capacity {
		b.resize(b.capacity * 2)
	}

	//indicar o início
	b.front = (b.front - 1 + b.capacity) % b.capacity
	b.data[b.front] = value
	b.size++
}

// popback para próximo teste
// PopBack()            // remove valor do fim
// semelhante ao popfront
func (b *Deque) PopBack() error {
	//verificação do tamanho
	if b.size == 0 {
		return fmt.Errorf("fail: buffer vazio")
	}

	//>1,2,3_
	//size= 3
	//popback
	//>1,2_ _
	//size = 2
	//indicar o fim
	//(front + size -1) % capacity
	//pode só reduzir o size tbm
	b.size--
	return nil

}

//agora é front e back

func (b *Deque) Front() (int, error) { //retorna valor do início
	//verificar se tem size
	if b.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}

	return b.data[b.front], nil

}

func (b *Deque) Back() (int, error) { //retorna valor do final
	//verificar se tem size
	if b.size == 0 {
		return 0, fmt.Errorf("fail: buffer vazio")
	}
	//indicar o final
	index := (b.size - 1 + b.front) % b.capacity
	return b.data[index], nil

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := &Deque{data: make([]int, 4), capacity: 4}

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
			fmt.Println(buf.String())
		case "debug":
			fmt.Println(buf.Debug())
		case "size":
			fmt.Println(buf.Len())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				buf.PushFront(num)
			}
		case "pop_back":
			if err := buf.PopBack(); err != nil {
				fmt.Println(err)
			}
		case "pop_front":
			if err := buf.PopFront(); err != nil {
				fmt.Println(err)
			}
		case "front":
			if val, err := buf.Front(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "back":
			if val, err := buf.Back(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		case "clear":
			buf.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
