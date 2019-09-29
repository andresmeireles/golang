package main

import "fmt"

func main() {
	//criando maps vazios em go
	mapOfPeople := make(map[int]string)
	/*
	* o exemplo acima cria um map com chave com valor int
	* que receberá valores do tipo string
	 */

	mapOfPeople[1] = "José"
	// mapOfPeople["fred"] = 82 <- Esse exemplo vai dar errado pois os tipos estão incorretos

	fmt.Printf("seu nome é %v", mapOfPeople[1])

	for nome, id := range mapOfPeople {
		fmt.Print("\n", nome, " ", id)
	}

	// removeção de apps
	delete(mapOfPeople, 1)
}
