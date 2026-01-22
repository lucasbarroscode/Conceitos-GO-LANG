package main

import "fmt"

//o fato de usar o ... indica que pode receber parametros variavies
func media(numeros ...float64) float64 {
	total := 0.0
	if numeros != nil {
		for _, num := range numeros {
			total += num
		}
		return total / float64(len(numeros))
	}
	return 0

}

func main() {
	fmt.Printf("Média: %.2f", media())
}
