package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	qnt      float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	// usando structs
	var produto1 = produto{
		nome:     "Nektok",
		preco:    250,
		qnt:      4,
		desconto: 15,
	}

	fmt.Printf("Nome do produto %s, preço %f, disponivel %f, com desconto de %f \n", produto1.nome, produto1.preco, produto1.qnt, produto1.desconto)

	// criando struct em sintaxe curta
	produto2 := produto{"Nektok", 300, 200, 4}

	fmt.Printf("Nome do produto %s, preço %f, disponivel %f, com desconto de %f \n", produto2.nome, produto2.preco, produto2.qnt, produto2.desconto)
}
