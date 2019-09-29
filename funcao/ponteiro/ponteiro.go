package main

import "fmt"

func inc1(n int) {
	n++
}

// marcação do ponteiro é no tipo da variavel
func inc2(n *int) {
	// por um asterisco [*] em uma variavel a
	// desreferencia, ou seja, pega o valor da
	// referencia
	*n++
	// neste caso a varaivel que é uma referencia
	// da memoria teve sua referencia acessada e
	// seu valor retornado
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)

}
