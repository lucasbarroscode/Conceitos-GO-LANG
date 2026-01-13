package main

//é possivel nomear o import, por exemplo abaixo nomeiei o import como m e chamei m la em baixo

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("Area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)
}
