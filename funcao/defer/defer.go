package main

import "fmt"

func valorDaConta(valor int) int {
	/**
	* defer faz o que esta após ele ser executado no momento antes do
	* retorno de uma função
	 */
	defer fmt.Println("o fim")

	if valor > 5000 {
		fmt.Println("Valor alto")
		return 5000
	}
	fmt.Println("Valor baixo")
	return valor
}

func main() {
	fmt.Println(valorDaConta(3000))
}
