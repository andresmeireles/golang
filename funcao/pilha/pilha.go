package main

import "runtime/debug"

func f1() {
	debug.PrintStack() // mostra as funções chamadas até a função ser chamada
}

func f2() {
	f1()
}

func f3() {
	f2()
}

func f4() {
	f3()
}

func main() {
	f4()
}
