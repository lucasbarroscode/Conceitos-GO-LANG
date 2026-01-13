package main

import "fmt"

func main() {
	//se remover o [ ...] não vai mais ser um array e sim um slice
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	//o range retorna dois elementos o i=indice e elemento do array que está sendo percorrido
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	//_ ignorando o valor do indice
	for _, num := range numeros {
		fmt.Println(num)
	}
}
