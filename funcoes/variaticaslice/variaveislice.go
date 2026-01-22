package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	//é um slice pois não foi definido o tamanho
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}
