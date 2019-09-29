package main

import "fmt"

func funcWithReturn() string {
	return "Value returned"
}

// funções que não retornam nenhum valor não podem ter nenhum valor como retorno
func funcWithAction() {
	fmt.Println("Funcion with no return, this func has no return statement")
}

func funcWithParameters(p1 string, p2 float64) {
	fmt.Println("Nome ", p1, " valor float ", p2)
}

func funcWithParametersWithSameType(p1, p2 string) string {
	return fmt.Sprintf("Parameters %s and %s", p1, p2)
}

// funções com multiplos retornos ainda devem ter apenas uma declaração return
func funcWithMultipleReturns() (string, string, string) {
	return "string1", "string2", "String3"
}

func main() {
	fmt.Println("começando a usar as funções")

	fmt.Println("Função com retorno ", funcWithReturn())

	funcWithAction()

	funcWithParameters("Parametro 1", 800)

	funcWithParametersWithSameType("Parametro string ", "Parametro string 2")

	// ao usar variaveis para como retorno, todas as variaveis precisam tratadas, se
	// caso uma função com três retornos só tem dois atributos então o programa
	// retornará erro.
	r1, r2, _ := funcWithMultipleReturns()

	fmt.Println(r1)
	fmt.Println(r2)
}
