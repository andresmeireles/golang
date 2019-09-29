package main

import "fmt"

type item struct {
	produtoID  int
	quantidade float64
	preco      float64
}

type pedido struct {
	pedidoID int
	items    []item
}

func (p pedido) valorDoPedido() float64 {
	total := 0.0
	for _, itemPrice := range p.items {
		total += (itemPrice.preco * itemPrice.quantidade)
	}

	return total
}

func main() {
	item1 := item{82, 3, 450}
	item2 := item{81, 4, 150}
	item3 := item{83, 2, 650}
	item4 := item{84, 1, 250}

	pedido1 := pedido{
		pedidoID: 63,
		items: []item{
			item1,
			item2,
			item3,
			item4,
		},
	}

	total := pedido1.valorDoPedido()

	fmt.Printf("%f", total)

}
