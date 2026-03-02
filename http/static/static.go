package main

import (
	"log"
	"net/http"
)

// so carrega no brownser se rodar via linha de comando
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
