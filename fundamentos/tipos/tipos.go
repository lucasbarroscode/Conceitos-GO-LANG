package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numero inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//sem sinal ( só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal
	il := math.MaxInt64
	fmt.Println("Literal inteiro é", il)
	fmt.Println("O byte é", reflect.TypeOf(il))

	//numeros reais (float 32 e float64)

	var x = float32(49.99)
	fmt.Println("Tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99)) //float 64

	//boolean
	bo := true
	fmt.Println("Tipo de BO é ", reflect.TypeOf(bo))
	fmt.Print(!bo)

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Leo`
	fmt.Println("O tamanho da string é", len(s2))

}
