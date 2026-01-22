package main

import "fmt"

//é possivel ter um init para cada arquivo.
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
