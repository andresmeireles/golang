package main

import "fmt"

func imprimirAprovados(nomes ...string) {
	fmt.Println("Lista dos aprovados")
	for i, nome := range nomes {
		fmt.Printf("%d) %s\n", i, nome)
	}
}

func main() {
	aprovados := []string{"Andre", "Yasmim", "Fabricio"}

	// função permite apenas slices
	imprimirAprovados(aprovados...)
}
