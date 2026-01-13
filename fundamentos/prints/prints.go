package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	fmt.Println("nova")
	fmt.Println("linha")

	x := 3.1415

	//Não é possivel fazer dessa forma
	//fmt.Println("o valoir de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("o valoir de x é ", xs)
	fmt.Println("o valoir de x é ", x)

	fmt.Printf("o valoir de x é %.2f", x)
}
