package main

import "fmt"

func multiplicar(n1 int, n2 int) int {
	return n1 * n2
}

func wrongFunc(p1, p2 int) int {
	return p1
}

// ao receber uma função como valor, a função precisa ter a assinatura
// desejada pela funçao
func exec(funcToExec func(int, int) int, p1, p2 int) int {
	return funcToExec(p1, p2)
}

func main() {
	fmt.Println(exec(wrongFunc, 2, 2))
}
