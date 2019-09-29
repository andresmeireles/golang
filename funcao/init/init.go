package main

import "fmt"

// a função init é rodada antes de toddas as funções
func init() {
	fmt.Println("Esse texto veio primeiro que a função main")
}

func main() {
	fmt.Println("Texto teste")
}
