package main

import "fmt"

func main() {
	// criado maps com valores
	// os valores são inseridos com de forma semelhante a json
	funcionarioSalario := map[string]float64{
		"Andre":    1400.00,
		"Fabricio": 1700.00,
		"Debora":   0.00, // maps precisam terminar em ,
	}

	// atribuir valores não poribe atribuição após criação
	funcionarioSalario["Simone"] = 40000

	// removeção de itens inexistentes não gera erro
	delete(funcionarioSalario, "zeca")

	for nome, salario := range funcionarioSalario {
		fmt.Println("nome do meliante ", nome, " salario do meliante ", salario)
	}
}
