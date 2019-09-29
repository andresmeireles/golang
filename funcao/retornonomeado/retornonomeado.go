package main

import "fmt"

func namedReturn(p1, p2 int) (primeiro int, segundo int) /* mesmas regras valem para declaração de retornos */ {
	primeiro = p1
	segundo = p2
	return // retorno limpo. No caso não precisa definir os parametros retornados de forma explicita
}

func noNamedReturn(p1, p2 int) (int, int) {
	return p1, p2
}

func main() {
	v1, v2 := namedReturn(2, 3)

	fmt.Println(v1, v2)

	v3, v4 := noNamedReturn(4, 5)

	fmt.Println(v3, v4)
}
