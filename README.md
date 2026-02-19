## Compilar
Por conta da extensão CODE RUNNER, compilar o projeto com o comando CTRL + ALT + N <br>
E para parar a execução de um loop CTRL + ALT + M
<br>
##Configuração
Na versão 1.11 do Go, foi introduzido um sistema de módulos na linguagem que é incompatível com as próximas aulas. Caso você esteja utilizando a versão 1.11 ou superior, <br>
você precisará rodar o comando GO111MODULE=off para desativar este sistema e poder acompanhar as próximas aulas.
<br>
<h3>Paralelismo: Executar codigo simultanemante em processadores fisicos diferentes.
<br>
<h3>Concorrencia: Intercalar (adminsitrar) varios processos ao mesmo tempo e isso pode ocorrer em um único processador fisico.
<br>
##Teste
Para gerar o resultado dentro desse arquivo: 
go test --coverprofile=resultado.out
<br>
Conseguir ler esse resultado e mostrar na tela o que está acontecendo: go tool cover -gunc=resultado.out
<br>
Pagina html com o resultado detalhado com a sua cobertura: go tool cover -html=resultado.out