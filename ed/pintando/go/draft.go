package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	var s, area, valor float64
	s = (a + b + c) / 2
	valor = s * (s - a) * (s - b) * (s - c)
	area = math.Sqrt(valor)
	fmt.Printf("%.2f\n", area)

}
