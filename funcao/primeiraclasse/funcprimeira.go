package main

import "fmt"

func main() {
	// função anônima
	funcAnon := func(b1 string) string {
		return b1
	}

	fmt.Println(funcAnon("bola"))
}
