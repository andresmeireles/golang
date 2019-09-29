package main

import "fmt"

func closure() func() {
	x := 10

	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := 20

	fmt.Println(x)

	pX := closure()

	pX()
}
