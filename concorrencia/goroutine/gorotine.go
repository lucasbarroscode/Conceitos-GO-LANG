package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de vc!", 1)

	//nesse caso nao ira logar, pois a aplicacao acabou antes de dar um segundo
	//se eu forcar ela esperar um pouco.. igual ali embaixo, ela ira executar por 5 segundo, mas depois ira se encerrar!
	//funcoes sendo chamadas de forma independente
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	//time.Sleep(time.Second * 5)
	//fmt.Println("Fim")

	//inserir o go antes do metodo indica que ele ira executar uma goroutine
	//o codigo ira terminar qanddo o joao falar 5 vezes, mesmo maria nao terminando as suas 10 vezes
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
}
