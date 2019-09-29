package main

import "fmt"

/*
* Para simplificar esse codigo poderiamos usar inteiros unsigned
* pois como eles não tem sinal, não seria possível passar numeros
* negativos, a propria linguagem já validaria o codigo.
 */
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número invalido %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	positivo, e := fatorial(1)
	negativo, x := fatorial(-1)
	inteiro, error2 := fatorial(9)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(positivo)
	}
	if x != nil {
		fmt.Println(x)
	} else {
		fmt.Println(negativo)

	}
	if error2 != nil {
		fmt.Println(error2)
	} else {
		fmt.Println(inteiro)
	}
}
