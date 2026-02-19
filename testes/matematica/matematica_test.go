package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

//para executar aperte esse run test e la em cima para rodar todos os testes do pacote
//ou via terminal (dentro da pasta) go test
//precisa comecar com o prefixo de Test os metodos de teste
func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
